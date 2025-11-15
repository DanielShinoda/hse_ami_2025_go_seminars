package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go worker(ctx, "Deadline Worker")

	if dl, ok := ctx.Deadline(); ok {
		fmt.Printf("Дедлайн установлен на: %v\n", dl)
	}

	time.Sleep(3 * time.Second)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: получил сигнал отмены, завершаю работу\n", name)
			return
		default:
			fmt.Printf("%s: работаю...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
