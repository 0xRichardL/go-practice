package main

import (
	"fmt"
	"time"
)

func work(ch chan<- int) {
	// Change the sleep duration to test timeout behavior.
	time.Sleep(3 * time.Second)
	ch <- 42
}

func main() {
	ch := make(chan int)
	fmt.Println("Start at:", time.Now().String())
	go work(ch)
	select {
	case value := <-ch:
		fmt.Println("Received value:", value)
	case t := <-time.After(2 * time.Second):
		fmt.Println("Timeout! at:", t.String())
	}
}
