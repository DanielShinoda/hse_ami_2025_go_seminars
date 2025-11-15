package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"example/src/context_http/server/handlers"
	"example/src/context_http/server/middleware"
)

func main() {
	r := mux.NewRouter()

	// Инициализируем обработчики
	deckHandler := handlers.NewDeckHandler()

	// API routes с версионированием
	api := r.PathPrefix("/api/v1").Subrouter()

	// Deck routes
	api.HandleFunc("/decks", deckHandler.CreateDeck).Methods("POST")
	api.HandleFunc("/decks/{deckId}", deckHandler.GetDeck).Methods("GET")
	api.HandleFunc("/decks/{deckId}/draw", deckHandler.DrawCards).Methods("POST")
	api.HandleFunc("/decks/{deckId}/shuffle", deckHandler.ShuffleDeck).Methods("POST")

	// Swagger документация
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "healthy", "timestamp": "` + time.Now().Format(time.RFC3339) + `"}`))
	}).Methods("GET")

	// Применяем middleware ко всем роутам
	wrappedRouter := middleware.ApplyMiddleware(r)

	// Настройки сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      wrappedRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("🚀 Server started on http://localhost:8080")
	log.Println("📚 Swagger docs available at http://localhost:8080/swagger/")
	log.Println("❤️  Health check at http://localhost:8080/health")

	log.Fatal(server.ListenAndServe())
}
