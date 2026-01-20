package main

import (
	"fmt"
	"sync"
)

func main() {
	once := &sync.Once{}
	config := func(value int) {
		fmt.Println("Configured:", value)
	}
	for i := range 5 {
		go func() {
			once.Do(func() {
				config(i)
			})
		}()
	}
}
