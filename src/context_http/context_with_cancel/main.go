package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "Worker 1")

	time.Sleep(2 * time.Second)

	fmt.Println("Main: отправляю сигнал отмены")
	cancel()

	time.Sleep(1 * time.Second)
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
