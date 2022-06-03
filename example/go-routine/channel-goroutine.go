package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// generate naturals numbers 0 to 100
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// used naturals number & calculate Squarer of it
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}