package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	counter := 0
	// Function to increment the counter safely
	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		local := counter
		time.Sleep(10 * time.Millisecond) // Simulate some work
		local++
		counter = local
	}

	var wg sync.WaitGroup
	// Start multiple goroutines to increment the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter)
}
