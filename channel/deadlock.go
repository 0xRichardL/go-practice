package main

import (
	"fmt"
	"time"
)

func direct_deadlock() {
	ch := make(chan int)
	// Attempting to send without a corresponding receiver leads to deadlock.
	// At this point, the program will check the number of goroutines and find none can proceed.
	// Hence, it will panic with: "fatal error: all goroutines are asleep - deadlock!"
	ch <- 1
	// This line is never reached
	value := <-ch
	fmt.Println("Received:", value)
}

func dodged_deadlock_1() {
	ch := make(chan int)
	// Start a goroutine to receive from the channel first.
	go func() {
		value := <-ch
		fmt.Println("Received:", value)
	}()
	ch <- 1
	// Wait for printing in the goroutine.
	time.Sleep(1 * time.Second)
}

func dodged_deadlock_2() {
	ch := make(chan int)
	// Separate the producer in other goroutine.
	go func() {
		ch <- 1
	}()
	value := <-ch
	fmt.Println("Received:", value)
}

func buffered_no_deadlock() {
	ch := make(chan int, 1)
	ch <- 1
	value := <-ch
	fmt.Println("Received:", value)
}

func buffered_deadlock() {
	ch := make(chan int, 1)
	ch <- 1
	// Deadlock occurs here because the buffer is full and there's no receiver.
	ch <- 2
	value := <-ch
	fmt.Println("Received:", value)
}

func no_sender_deadlock() {
	ch := make(chan int)
	// Spawn a goroutine that sleeps for a while before exiting.
	go func() {
		time.Sleep(10 * time.Second)
	}()
	// This will delay the deadlock detection by 10 seconds.
	// When the delay routine exits, the program will detect no other goroutines can proceed.
	value := <-ch
	fmt.Println("Received:", value)
}

func main() {
	// direct_deadlock()
	// dodged_deadlock_1()
	// dodged_deadlock_2()
	// buffered_no_deadlock()
	// buffered_deadlock()
	// no_sender_deadlock()
}
