package main

import (
	"fmt"
	"time"
)

// Single goroutine for sending and main goroutine receiving from a channel.
func single() {
	ch := make(chan string)

	go func() {
		// Send a message to the channel.
		ch <- "Hello from goroutine"
	}()
	// Receive the message from the channel.
	// This will block until a message is available.
	msg := <-ch
	fmt.Printf("Received from goroutine: \"%s\"\n", msg)
	close(ch)
	msg, ok := <-ch // Receiving from a closed channel returns the zero value.
	fmt.Printf("Received from closed channel: \"%s\", ok: %t\n", msg, ok)
}

func multiple() {
	ch := make(chan string)

	// Start multiple sender goroutines.
	for i := 0; i < 3; i++ {
		go func(id int) {
			now := time.Now().Nanosecond()
			msg := fmt.Sprintf("Hello from goroutine %d at %dns", id, now)
			ch <- msg
			// Case 2: Safe close(ch) at last process
			if i == 2 {
				close(ch)
			}
		}(i)
		// This is needed to avoid race condition when spawning close goroutines.
		// Millis is safe thread.
		time.Sleep(1 * time.Millisecond)
		// These have a change to panic since the scheduled duration too small.
		time.Sleep(1 * time.Microsecond)
		time.Sleep(1 * time.Nanosecond)
	}
	// Case 1 (safe): Only read a expected messages.
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Printf("Received: \"%s\"\n", msg)
	}
	// Case 2: Will read all message but deadlock at the end since no close(ch) from the last sender.
	for msg := range ch {
		fmt.Printf("Received: \"%s\"\n", msg)
	}
}

func main() {
	single()
	multiple()
}
