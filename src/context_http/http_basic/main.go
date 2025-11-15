package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"example/src/context_http/client"
	"example/src/context_http/types"
)

func main() {
	main2()
}

func main1() {
	deck, err := client.CreateNewDeck()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Создана колода: %s\n", deck.DeckID)
	fmt.Printf("Карт в колоде: %d\n", deck.Remaining)

	draw, err := drawCards(deck.DeckID, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nВытянули %d карт:\n", len(draw.Cards))
	for i, card := range draw.Cards {
		fmt.Printf("%d. %s %s (%s)\n", i+1, card.Value, card.Suit, card.Code)
	}
	fmt.Printf("Осталось карт: %d\n", draw.Remaining)
}

func drawCards(deckID string, count int) (*types.DrawResponse, error) {
	url := fmt.Sprintf("https://deckofcardsapi.com/api/deck/%s/draw/?count=%d", deckID, count)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var draw types.DrawResponse
	err = json.Unmarshal(body, &draw)
	if err != nil {
		return nil, err
	}

	return &draw, nil
}

func main2() {
	api := client.NewCardAPI()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deck, err := api.CreateDeckWithContext(ctx, true, 1)
	if err != nil {
		fmt.Println("Error creating deck:", err)
		return
	}

	fmt.Printf("Создана колода: %s\n", deck.DeckID)

	for i := 0; i < 3; i++ {
		draw, err := api.DrawCardsWithContext(ctx, deck.DeckID, 2)
		if err != nil {
			fmt.Println("Error drawing cards:", err)
			return
		}

		fmt.Printf("\nРука %d:\n", i+1)
		for j, card := range draw.Cards {
			fmt.Printf("  %d. %s of %s\n", j+1, card.Value, card.Suit)
		}
		fmt.Printf("Осталось карт: %d\n", draw.Remaining)

		time.Sleep(500 * time.Millisecond) // Небольшая задержка
	}
}
