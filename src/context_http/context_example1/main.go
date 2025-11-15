package main

import (
	"context"
	"fmt"
	"time"
)

type key string

const traceIDKey key = "trace_id"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, traceIDKey, "trace-abc-123")

	go complexWorker(ctx)

	time.Sleep(3 * time.Second)
}

func complexWorker(ctx context.Context) {
	traceID := ctx.Value(traceIDKey)

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Printf("Worker (Trace: %v): остановлен. Причина: %v\n",
				traceID, err)
			return
		default:
			fmt.Printf("Worker (Trace: %v): итерация %d\n", traceID, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
