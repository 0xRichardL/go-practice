package main

import (
	"fmt"
	"time"
)

func throttle(rate int, ttl time.Duration) chan time.Time {
	// the limiter use a full buffered channel to hold tokens.
	limiter := make(chan time.Time, rate)
	// fill the channel to capacity.
	for i := 0; i < rate; i++ {
		limiter <- time.Now()
	}
	go func() {
		// refill the channel at the specified rate.
		for range time.Tick(ttl) {

			for i := 0; i < rate; i++ {
				// try to add a token, but don't block if the channel is full.
				select {
				case limiter <- time.Now():
				default:
					i = rate
				}
			}
		}
	}()

	return limiter
}

func main() {
	// Allow 5 requests per second
	limit := throttle(5, 1*time.Second)

	for i := 0; i < 20; i++ {
		<-limit // take token: blocks if rate exceeded
		if i == 3 || i == 13 || i == 17 {
			time.Sleep(3 * time.Second) // simulate a pause in requests
		}
		fmt.Println("Request", i, "at", time.Now().Format("15:04:05.000"))
	}
}
