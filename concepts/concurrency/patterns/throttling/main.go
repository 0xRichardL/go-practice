package main

import (
	"fmt"
	"sync"
	"time"
)

func throttle(rate int, ttl time.Duration) (<-chan struct{}, func()) {
	// Use a buffered channel as a token bucket. struct{} carries no data.
	limiter := make(chan struct{}, rate)
	// Start with a full bucket, allowing an initial burst of up to rate operations.
	for range rate {
		limiter <- struct{}{}
	}

	done := make(chan struct{})
	var cancelOnce sync.Once
	cancel := func() {
		cancelOnce.Do(func() {
			close(done)
		})
	}

	go func() {
		ticker := time.NewTicker(ttl)
		defer ticker.Stop()
		// Refill the bucket to capacity at each interval.
		for {
			// Wait for cancellation or the next refill interval.
			select {
			case <-done:
				return

			case <-ticker.C:
			refill:
				for {
					// Add tokens until the bucket is full without blocking.
					select {
					case limiter <- struct{}{}:
					default:
						break refill
					}
				}
			}
		}
	}()

	return limiter, cancel
}

func main() {
	// Allow bursts of up to five requests and refill the bucket every second.
	limit, cancel := throttle(5, 1*time.Second)
	defer cancel()
	for i := range 20 {
		<-limit // Consume one token, blocking while the bucket is empty.
		fmt.Println("Request", i, "at", time.Now().Format("15:04:05.000"))
		if i == 3 || i == 13 || i == 17 {
			time.Sleep(3 * time.Second) // Simulate idle time, allowing the bucket to refill.
		}
	}
}
