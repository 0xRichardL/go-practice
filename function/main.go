package main

import "fmt"

// Named return values.
// Tuple returns.
func namedReturnValues() (x int, y int) {
	x = 5
	y = 10
	return // naked return
}

// Variadic function.
// the ...type should be at last parameter.
func variadicFn(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

// Higher-order function example.
func higherOrderFn(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

func main() {
	x, y := namedReturnValues()
	fmt.Println(x, y)

	// Anonymous function assigned to a variable.
	anonymousFunc := func(a int, b int) int {
		return a + b
	}
	fmt.Println("Anonymous function:", anonymousFunc(3, 4))

	fmt.Println("Variadic function sum:", variadicFn(1, 2, 3, 4, 5))

	fmt.Println("Higher-order function:", higherOrderFn(
		func(a int, b int) int { return a * b },
		3,
		4,
	))
}
