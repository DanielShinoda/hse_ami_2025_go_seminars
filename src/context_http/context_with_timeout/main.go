package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx, "Timeout Worker")

	time.Sleep(5 * time.Second)
	fmt.Println("Main: завершение")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Printf("%s: завершение. Причина: %v\n", name, err)
			return
		default:
			fmt.Printf("%s: выполняю задачу...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}
