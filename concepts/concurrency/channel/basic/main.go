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

func try_receive() {
	test := func() bool {
		ch := make(chan int)
		// Spawn a goroutine to send a value aim for a race condition.
		go func() {
			ch <- 1
		}()
		time.Sleep(230 * time.Nanosecond) // Optimized for 50:50 chances.
		select {
		case <-ch:
			return true
		default:
			return false
		}
	}

	catches := 0
	misses := 0
	for i := 0; i < 1000; i++ {
		res := test()
		if res {
			catches++
		} else {
			misses++
		}
	}
	fmt.Printf("Catch: %d, Missed: %d\n", catches, misses)
}

func main() {
	// single()
	// multiple()
	// try_receive()
}
