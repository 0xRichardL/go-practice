package main

import (
	"fmt"
	"time"
)

func unsafe() {
	// Create an unbuffered channel.
	ch := make(chan string)
	go func() {
		close(ch)
	}()
	time.Sleep(1 * time.Second)
	// Send message to a closed channel will cause panic.
	ch <- "Hello"
}

// Multiple goroutines may close the channel concurrently.
// It use a done channel to signal when 1 goroutine will close the main channel at next step.
// A very tiny chance that done signal arrives after the send operation, still may cause panic.
func lowRisk() {
	ch := make(chan string)
	done := make(chan bool)

	// This goroutine will close the channel after a delay.
	go func() {
		time.Sleep(1 * time.Second)
		done <- true
		close(ch)
	}()
	// This goroutine will attempt to receive messages from the channel.
	go func() {
		time.Sleep(1 * time.Second)
		for v := range ch {
			fmt.Printf("Received: \"%s\"\n", v)
		}
	}()
	// Depend on which goroutine runs first, we may send or find the channel closed.
	select {
	case ch <- "Hello":
		fmt.Println("Sent message to channel")
	case <-done:
		fmt.Println("Channel closed before sending message")
	}
}

func main() {
	unsafe()
	lowRisk()
}
