package main

import "fmt"

func recoverFn() {
	if r := recover(); r != nil {
		println("Recovered from panic:", r.(string))
	}
}

func deferFn() {
	println("Defer function executed.")
}

func panicFn() {
	// Recover should placed in the service function not main function.
	defer recoverFn()
	// Deferred function still work well.
	defer deferFn()
	panic("Panic triggered!")
	// Never reached since panic makes control flow jump to end of closure.
	fmt.Println("Never reached here.")
}

func routinePanic() {
	// Not working cuz this on main goroutine.
	// Not in the goroutine where panic happens.
	defer recoverFn()
	go func() {
		// recover works here.
		defer recoverFn()
		panic("Panic in goroutine!")
	}()
}

func main() {
	// panicFn()
	routinePanic()
	fmt.Println("Continue without panic.")
}
