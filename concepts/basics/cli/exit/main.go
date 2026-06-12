package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("Started at: %v\n", time.Now())
	// This deferred function will not run since os.Exit is called
	defer func() {
		fmt.Println("Deferred function never reached")
	}()
	// Stop the program from a goroutine
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("Finished at: %v\n", time.Now())
		os.Exit(1)
	}()

	// Sleep for a while
	time.Sleep(5 * time.Second)
	fmt.Println("This line will not be printed")
}
