package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"example/src/context_http/types"
)

type CardAPI struct {
	BaseURL string
	Client  *http.Client
}

func NewCardAPI() *CardAPI {
	return &CardAPI{
		BaseURL: "https://deckofcardsapi.com/api",
		Client: &http.Client{
			Timeout: 10 * time.Second, // Таймаут на весь запрос
		},
	}
}

func (api *CardAPI) CreateDeckWithContext(ctx context.Context, shuffled bool, deckCount int) (*types.DeckResponse, error) {
	url := fmt.Sprintf("%s/deck/new/", api.BaseURL)
	if shuffled {
		url += "shuffle/"
	}

	// Добавляем параметры
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if deckCount > 1 {
		q.Add("deck_count", fmt.Sprintf("%d", deckCount))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deck types.DeckResponse
	err = json.NewDecoder(resp.Body).Decode(&deck)
	if err != nil {
		return nil, err
	}

	return &deck, nil
}

func (api *CardAPI) DrawCardsWithContext(ctx context.Context, deckID string, count int) (*types.DrawResponse, error) {
	url := fmt.Sprintf("%s/deck/%s/draw/", api.BaseURL, deckID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("count", fmt.Sprintf("%d", count))
	req.URL.RawQuery = q.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var draw types.DrawResponse
	err = json.NewDecoder(resp.Body).Decode(&draw)
	if err != nil {
		return nil, err
	}

	return &draw, nil
}

func CreateNewDeck() (*types.DeckResponse, error) {
	resp, err := http.Get("https://deckofcardsapi.com/api/deck/new/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var deck types.DeckResponse
	err = json.Unmarshal(body, &deck)
	if err != nil {
		return nil, err
	}

	return &deck, nil
}
