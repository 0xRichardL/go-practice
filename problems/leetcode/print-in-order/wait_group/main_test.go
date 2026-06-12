package main

import "testing"

func Benchmark_main(b *testing.B) {
	for b.Loop() {
		main()
	}
}
