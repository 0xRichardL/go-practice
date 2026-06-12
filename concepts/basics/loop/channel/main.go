package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	// 2 goroutines reading from the same channel.
	go func() {
		// The v will be received only in 1 of the goroutines.
		for v := range ch {
			fmt.Println("Goroutine #1: ", v)
		}
	}()
	go func() {
		for v := range ch {
			fmt.Println("Goroutine #2: ", v)
		}
	}()
	// Give 100 values to channel and then close it
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
