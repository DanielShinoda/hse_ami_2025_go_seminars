package domain

import "time"

// Card представляет игральную карту
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// Deck представляет колоду карт
type Deck struct {
	ID        string    `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []Card    `json:"cards,omitempty"`
	CreatedAt time.Time `json:"-"`
}

// CreateDeckRequest запрос на создание колоды
type CreateDeckRequest struct {
	Shuffled bool     `json:"shuffled" example:"false"`
	Cards    []string `json:"cards,omitempty" example:"AS,KD,QH,JC"`
}

// CreateDeckResponse ответ при создании колоды
type CreateDeckResponse struct {
	DeckID    string `json:"deck_id" example:"a251071b-662f-44b6-ba11-e24863039c59"`
	Shuffled  bool   `json:"shuffled" example:"false"`
	Remaining int    `json:"remaining" example:"52"`
}

// DrawCardsResponse ответ при взятии карт
type DrawCardsResponse struct {
	Cards []Card `json:"cards"`
}

// ErrorResponse ответ с ошибкой
type ErrorResponse struct {
	Error string `json:"error" example:"deck not found"`
}
