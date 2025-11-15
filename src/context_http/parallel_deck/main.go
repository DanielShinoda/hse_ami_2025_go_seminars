package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"example/src/context_http/client"
	"example/src/context_http/types"
)

func dealCardsToPlayers(ctx context.Context, deckID string, players int, cardsPerPlayer int) (map[int][]types.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	results := make(chan types.PlayerResult, players)
	var wg sync.WaitGroup

	for i := 1; i <= players; i++ {
		wg.Add(1)
		go func(playerID int) {
			defer wg.Done()

			start := time.Now()
			cards, err := drawCardsWithRetry(ctx, deckID, cardsPerPlayer, 3)
			duration := time.Since(start)

			results <- types.PlayerResult{
				PlayerID: playerID,
				Cards:    cards,
				Error:    err,
				Duration: duration,
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Собираем результаты
	playerCards := make(map[int][]types.Card)
	var errors []error

	for result := range results {
		if result.Error != nil {
			errors = append(errors, fmt.Errorf("игрок %d: %v", result.PlayerID, result.Error))
			log.Printf("❌ Игрок %d: ошибка за %.2fс: %v",
				result.PlayerID, result.Duration.Seconds(), result.Error)
		} else {
			playerCards[result.PlayerID] = result.Cards
			log.Printf("✅ Игрок %d: получил %d карт за %.2fс",
				result.PlayerID, len(result.Cards), result.Duration.Seconds())
		}
	}

	if len(errors) > 0 {
		return playerCards, fmt.Errorf("ошибки при раздаче: %v", errors)
	}

	return playerCards, nil
}

func drawCardsWithRetry(ctx context.Context, deckID string, count int, maxRetries int) ([]types.Card, error) {
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			cards, err := drawCards(ctx, deckID, count)
			if err == nil {
				return cards, nil
			}

			lastErr = err
			log.Printf("⚠️  Попытка %d/%d не удалась: %v", attempt, maxRetries, err)

			if attempt < maxRetries {
				delay := time.Duration(attempt*attempt) * 100 * time.Millisecond
				time.Sleep(delay)
			}
		}
	}

	return nil, fmt.Errorf("после %d попыток: %w", maxRetries, lastErr)
}

func drawCards(ctx context.Context, deckID string, count int) ([]types.Card, error) {
	url := fmt.Sprintf("https://deckofcardsapi.com/api/deck/%s/draw/?count=%d", deckID, count)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	var drawResponse types.DrawResponse
	if err := json.NewDecoder(resp.Body).Decode(&drawResponse); err != nil {
		return nil, err
	}

	if !drawResponse.Success {
		return nil, fmt.Errorf("API error: не удалось вытянуть карты")
	}

	return drawResponse.Cards, nil
}

func main() {
	ctx := context.Background()

	deck, err := client.CreateNewDeck()
	if err != nil {
		log.Fatal("Ошибка создания колоды:", err)
	}

	fmt.Printf("🎴 Создана колода: %s\n", deck.DeckID)
	fmt.Printf("📊 Карт в колоде: %d\n\n", deck.Remaining)

	players := 4
	cardsPerPlayer := 5

	fmt.Printf("🎮 Начинаем раздачу для %d игроков по %d карт...\n\n", players, cardsPerPlayer)

	playerCards, err := dealCardsToPlayers(ctx, deck.DeckID, players, cardsPerPlayer)
	if err != nil {
		log.Printf("⚠️  Были ошибки: %v", err)
	}

	fmt.Println("РЕЗУЛЬТАТЫ РАЗДАЧИ:")

	for playerID := 1; playerID <= players; playerID++ {
		cards := playerCards[playerID]
		fmt.Printf("\n👤 Игрок %d:\n", playerID)

		if len(cards) == 0 {
			fmt.Println("   ❌ Не получил карты")
			continue
		}

		for i, card := range cards {
			fmt.Printf("   %d. %s %s (%s)\n", i+1, card.Value, card.Suit, card.Code)
		}
	}

	remaining, err := checkRemainingCards(ctx, deck.DeckID)
	if err == nil {
		fmt.Printf("\n📊 Осталось карт в колоде: %d\n", remaining)
	}
}

func checkRemainingCards(ctx context.Context, deckID string) (int, error) {
	url := fmt.Sprintf("https://deckofcardsapi.com/api/deck/%s/", deckID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var deckResponse types.DeckResponse
	if err := json.NewDecoder(resp.Body).Decode(&deckResponse); err != nil {
		return 0, err
	}

	return deckResponse.Remaining, nil
}
