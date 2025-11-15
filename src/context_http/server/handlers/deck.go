package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"example/src/context_http/server/domain"
)

type DeckHandler struct {
	decks map[string]*domain.Deck
	mutex sync.RWMutex
}

func NewDeckHandler() *DeckHandler {
	return &DeckHandler{
		decks: make(map[string]*domain.Deck),
	}
}

// generateStandardCards создает стандартную колоду из 52 карт
func generateStandardCards() []domain.Card {
	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	var cards []domain.Card
	for _, suit := range suits {
		for _, value := range values {
			code := getCardCode(value, suit)
			cards = append(cards, domain.Card{
				Value: value,
				Suit:  suit,
				Code:  code,
			})
		}
	}
	return cards
}

// getCardCode генерирует код карты
func getCardCode(value, suit string) string {
	suitCode := string(suit[0])
	if value == "10" {
		return "0" + suitCode
	}
	valueCode := string(value[0])
	return valueCode + suitCode
}

// shuffleCards перемешивает карты
func shuffleCards(cards []domain.Card) []domain.Card {
	shuffled := make([]domain.Card, len(cards))
	copy(shuffled, cards)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

// CreateDeck создает новую колоду
func (h *DeckHandler) CreateDeck(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateDeckRequest

	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.sendError(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
	} else {
		shuffled := r.URL.Query().Get("shuffled")
		req.Shuffled = shuffled == "true"

		if cardsParam := r.URL.Query().Get("cards"); cardsParam != "" {
			for i := 0; i < len(cardsParam); i += 3 {
				end := i + 2
				if end > len(cardsParam) {
					end = len(cardsParam)
				}
				card := cardsParam[i:end]
				if len(card) == 2 {
					req.Cards = append(req.Cards, card)
				}
				if i+3 < len(cardsParam) && cardsParam[i+2] == ',' {
					i++
				}
			}
		}
	}

	deckID := fmt.Sprintf("%x", time.Now().UnixNano())

	var cards []domain.Card
	if len(req.Cards) > 0 {
		for _, code := range req.Cards {
			cards = append(cards, domain.Card{Code: code})
		}
	} else {
		cards = generateStandardCards()
	}

	// Перемешиваем если нужно
	if req.Shuffled {
		cards = shuffleCards(cards)
	}

	deck := &domain.Deck{
		ID:        deckID,
		Shuffled:  req.Shuffled,
		Remaining: len(cards),
		Cards:     cards,
		CreatedAt: time.Now(),
	}

	h.mutex.Lock()
	h.decks[deckID] = deck
	h.mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.CreateDeckResponse{
		DeckID:    deckID,
		Shuffled:  req.Shuffled,
		Remaining: len(cards),
	})
}

// GetDeck возвращает информацию о колоде
func (h *DeckHandler) GetDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deckID := vars["deckId"]

	h.mutex.RLock()
	deck, exists := h.decks[deckID]
	h.mutex.RUnlock()

	if !exists {
		h.sendError(w, "deck not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deck)
}

// DrawCards берет карты из колоды
func (h *DeckHandler) DrawCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deckID := vars["deckId"]

	countStr := r.URL.Query().Get("count")
	count := 1
	if countStr != "" {
		var err error
		count, err = strconv.Atoi(countStr)
		if err != nil || count < 1 {
			h.sendError(w, "invalid count parameter", http.StatusBadRequest)
			return
		}
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	deck, exists := h.decks[deckID]
	if !exists {
		h.sendError(w, "deck not found", http.StatusNotFound)
		return
	}

	if deck.Remaining < count {
		h.sendError(w, "not enough cards in deck", http.StatusBadRequest)
		return
	}

	drawnCards := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	deck.Remaining = len(deck.Cards)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.DrawCardsResponse{
		Cards: drawnCards,
	})
}

// ShuffleDeck перемешивает колоду
func (h *DeckHandler) ShuffleDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deckID := vars["deckId"]

	h.mutex.Lock()
	defer h.mutex.Unlock()

	deck, exists := h.decks[deckID]
	if !exists {
		h.sendError(w, "deck not found", http.StatusNotFound)
		return
	}

	deck.Cards = shuffleCards(deck.Cards)
	deck.Shuffled = true

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"deck_id": deckID,
	})
}

// sendError отправляет ошибку в формате JSON
func (h *DeckHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(domain.ErrorResponse{
		Error: message,
	})
}
