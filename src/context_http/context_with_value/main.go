package main

import (
	"context"
	"fmt"
)

type key string

const (
	userIDKey    key = "user_id"
	requestIDKey key = "request_id"
)

func main() {
	ctx := context.WithValue(context.Background(), userIDKey, "12345")
	ctx = context.WithValue(ctx, requestIDKey, "req-67890")

	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)

	fmt.Printf("Обрабатываем запрос:\n")
	fmt.Printf("  User ID: %v\n", userID)
	fmt.Printf("  Request ID: %v\n", requestID)

	if nonExistent := ctx.Value("non_existent_key"); nonExistent == nil {
		fmt.Printf("  Несуществующий ключ: %v\n", nonExistent)
	}
}
