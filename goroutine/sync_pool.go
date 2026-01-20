package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
)

// sync.Pool is a concurrent-safe object pool that stores temporary objects that can be reused.
// Use-case: Reuse expensive-to-create objects (like buffers) to reduce memory allocations
//
// Important notes:
// - Pool objects can be removed automatically during GC
// - Always reset object state before returning to pool
// - Don't store long-lived state in pooled objects

// Buffer Pool - most common use-case
// Demonstrates reusing byte buffers to avoid allocations in serialization/parsing
var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("  [Pool] Creating new buffer (pool was empty)")
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func putBuffer(buf *bytes.Buffer) {
	buf.Reset() // Important: clear the buffer before returning to pool
	bufferPool.Put(buf)
}

func processData(data string) string {
	buf := getBuffer()
	defer putBuffer(buf) // Always return to pool when done

	buf.WriteString("Processed: ")
	buf.WriteString(data)
	return buf.String()
}

func printGCStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("  [GC] NumGC: %d, Alloc: %v KB, TotalAlloc: %v KB\n",
		m.NumGC, m.Alloc/1024, m.TotalAlloc/1024)
}

func main() {
	fmt.Println("=== sync.Pool Example ===\n")

	// Initial GC stats
	fmt.Println("Initial state:")
	printGCStats()

	// Use buffer pool multiple times
	fmt.Println("\nProcessing data:")
	result1 := processData("Hello")
	result2 := processData("World")
	fmt.Printf("  %s\n  %s\n", result1, result2)

	// Check stats after processing
	fmt.Println("\nAfter processing:")
	printGCStats()

	// Force GC to see pool being cleared
	fmt.Println("\nForcing GC (this may clear the pool)...")
	runtime.GC()
	printGCStats()

	// Use again - might create new buffer if GC cleared the pool
	fmt.Println("\nProcessing after GC:")
	result3 := processData("Test")
	fmt.Printf("  %s\n", result3)
}
