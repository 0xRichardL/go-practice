package main

import (
	"fmt"
)

func f1() {
	boolMap := map[int]bool{
		1: true,
		2: false,
	}
	v1, ok1 := boolMap[1]
	fmt.Println(v1, ok1)
	v2, ok2 := boolMap[2]
	fmt.Println(v2, ok2)
	v3, ok3 := boolMap[3]
	fmt.Println(v3, ok3)
}

var intMap = map[int]int{}

type Data struct {
	A int
	B int
}

var structMap = map[int]Data{}

func main() {

}
