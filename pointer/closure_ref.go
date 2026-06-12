package main

import (
	"fmt"
	"time"
)

func main() {
	var fns []func()

	// BUGGY: manual loop doesn't get Go 1.22+ variable scoping fix
	fmt.Println("=== BUGGY VERSION (manual loop) ===")
	i := 0
	for i < 5 {
		fn := func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Buggy: i=%d\n", i)
		}
		fns = append(fns, fn)
		i++
	}
	// Execute all closures
	for _, fn := range fns {
		go fn()
	}
	time.Sleep(500 * time.Millisecond)

	// FIXED: capture by value
	fmt.Println("\n=== FIXED VERSION (capture by value) ===")
	fns = nil
	i = 0
	for i < 5 {
		i_copy := i // capture current value
		fn := func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Fixed: i=%d\n", i_copy)
		}
		fns = append(fns, fn)
		i++
	}
	// Execute all closures
	for _, fn := range fns {
		go fn()
	}
	time.Sleep(500 * time.Millisecond)
}
