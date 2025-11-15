package types

import "time"

type DeckResponse struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type DrawResponse struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Cards     []Card `json:"cards"`
	Remaining int    `json:"remaining"`
}

type Card struct {
	Code  string `json:"code"`
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Image string `json:"image"`
}

type PlayerResult struct {
	PlayerID int
	Cards    []Card
	Error    error
	Duration time.Duration
}
