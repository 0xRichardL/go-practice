package main

import "fmt"

func printAnything(a any) {
	fmt.Println(a)
}

func main() {
	var i any
	printAnything(i)
	i = 42
	printAnything(i)
	i = "hello"
	printAnything(i)
	i = struct{}{}
	printAnything(i)
	i = []int{1, 2, 3}
	printAnything(i)
}
