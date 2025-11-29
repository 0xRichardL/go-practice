package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a WaitGroup to wait for multiple goroutines to finish
	wg := sync.WaitGroup{}
	// Add tasks to the WaitGroup: Task 1 and Task 2
	wg.Add(2)
	go func() {
		// run wg.Done() when the goroutine completes.
		// Note: avoid to put this at the end, use defer.
		defer wg.Done()
		fmt.Println("Task 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Task 2")
	}()
	// Shortcut using a helper function to add and run a goroutine.
	// No manual wg.Add and wg.Done needed.
	wg.Go(func() {
		fmt.Println("Task 3")
	})
	// Wait for all goroutines to finish, tasks down to zero.
	// Main goroutine will be blocked here till all tasks are done.
	wg.Wait()
}
