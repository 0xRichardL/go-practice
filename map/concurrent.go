package main

import (
	"fmt"
	"sync"
)

// concurrent_map_error demonstrates the race condition when writing to a regular map concurrently
func concurrent_map_error() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m[n] = n * 2 // This will cause concurrent map writes panic
		}(i)
	}

	wg.Wait()
	fmt.Println("Regular map (if it didn't panic):", len(m))
}

func sync_map_basic() {
	var m sync.Map
	// Store values
	m.Store("a", 1)
	m.Store("b", 2)

	// Load values
	if val, ok := m.Load("a"); ok {
		fmt.Println("Key a:", val)
	}

	// Load or store value
	actual, loaded := m.LoadOrStore("c", 3)
	fmt.Println("Key c:", actual, "Loaded:", loaded)

	// Load the value then delete
	value, loaded := m.LoadAndDelete("b")
	fmt.Println("Deleted key b:", value, "Loaded:", loaded)

	// Swap value
	previous, swapped := m.Swap("b", 20)
	fmt.Println("Previous value for b:", previous, "Swapped:", swapped)

	// CompareAndSwap
	swapped = m.CompareAndSwap("a", 1, 10)
	fmt.Println("CompareAndSwap for a:", swapped)

	// CompareAndDelete
	deleted := m.CompareAndDelete("c", 3)
	fmt.Println("CompareAndDelete for c:", deleted)

	// Delete a key
	m.Delete("b")

	// Delete all entries
	m.Clear()

	// Range over map
	m.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true
	})
}

// sync_map_concurrent_write demonstrates the safe usage of sync.Map for concurrent operations
func sync_map_concurrent_write() {
	var m sync.Map
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Store(n, n*2)
		}(i)
	}

	wg.Wait()

	// Read values
	count := 0
	m.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	fmt.Println("sync.Map entries:", count)

	// Load a specific value
	if val, ok := m.Load(10); ok {
		fmt.Println("Value at key 10:", val)
	}
}

func main() {
	concurrent_map_error()
	sync_map_concurrent_write()
}
