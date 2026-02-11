package main

import (
	"runtime"
	"testing"
)

// Benchmark rapid small allocations
func BenchmarkRapidSmallAllocations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RapidSmallAllocations()
		b.StopTimer()
		runtime.GC()
		b.StartTimer()
	}
}
