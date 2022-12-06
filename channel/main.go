package main

import "fmt"

var a string = "1"
var ch = make(chan int, 1)

func f() {
	if a != "1" {
		return
	}
	a = "3"
	<-ch
}

func main() {
	go f()
	ch <- 0
	a = "2"
	fmt.Println(a)
}
