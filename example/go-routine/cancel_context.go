package main

import (
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range generateNumber(ctx) {
		fmt.Println(n)
		if n == 1000 {
			break
		}
	}

}

func generateNumber(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
