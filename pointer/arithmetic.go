package main

func main() {
	x := 42
	p := &x
	// Compile error: invalid operation: p + 1 (pointer arithmetic not allowed)
	p++
}
