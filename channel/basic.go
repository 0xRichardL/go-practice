package main

import "fmt"

func f() {
	fmt.Println("B")
}

func main() {
	// Trigger Go to spawn go Routine, instantly.
	// Create a new lightweight thread of execution and run this function independently.
	// Schedule the f() call at near future.
	go f()
	// f() can be scheduled here.
	fmt.Println("A")
	// f() can be scheduled here.
}

// f() can be scheduled here (when process exit and NO f() called).
