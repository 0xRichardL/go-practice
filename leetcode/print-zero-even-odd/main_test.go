package main

import "testing"

func Benchmark_main(t *testing.B) {
	for t.Loop() {
		main()
	}
}
