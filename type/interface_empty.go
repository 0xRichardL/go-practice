package main

import "fmt"

func printAnything(a interface{}) {
	fmt.Println(a)
}

func main() {
	// can hold any type
	var i interface{}
	// default value is nil
	printAnything(i)

	// assign int value
	i = 42
	// type assertion (unsafe)
	var j int = i.(int)
	printAnything(i)
	printAnything(j)

	// assign different types
	i = "hello"
	printAnything(i)
	i = struct{}{}
	printAnything(i)
	i = []int{1, 2, 3}
	printAnything(i)
}
