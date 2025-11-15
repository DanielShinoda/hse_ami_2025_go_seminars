package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.Background()
	ctx2 := context.TODO()

	fmt.Printf("Context 1: %v\n", ctx1)
	fmt.Printf("Context 2: %v\n", ctx2)
}
