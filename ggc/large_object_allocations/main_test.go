package main

import (
	"runtime"
	"testing"
)

// Benchmark large object allocations
func BenchmarkLargeObjectAllocations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargeObjectAllocations()
		b.StopTimer()
		runtime.GC()
		b.StartTimer()
	}
}
