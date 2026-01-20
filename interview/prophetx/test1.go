package main

import (
	"fmt"
	"sync"
)

/*

Write a program to print ProphetX 10 times, each character will be printed in each separate go routine 10 times.
Expected output:
ProphetX
ProphetX
.....
ProphetX

*/

func manager(topics []chan struct{}) {
	for range 10 {
		for _, topic := range topics {
			topic <- struct{}{}
			<-topic
		}
	}
	for _, topic := range topics {
		close(topic)
	}
}

func printChar(last bool, r rune, topic chan struct{}) {
	for true {
		_, ok := <-topic
		if !ok {
			return
		}
		fmt.Print(string(r))
		if last {
			fmt.Println()
		}
		topic <- struct{}{}
	}
}

func main() {
	wg := sync.WaitGroup{}
	text := "ProphetX"
	topics := make([]chan struct{}, len(text))
	for i := range topics {
		topics[i] = make(chan struct{})
	}
	go manager(topics)
	for i, c := range text {
		wg.Go(func() {
			printChar(i == len(text)-1, c, topics[i])
		})
	}
	wg.Wait()
}

/*
Concept: 2 types of routines: Controller, worker
- Worker: ready for any time and terminate when signal come.
- Controller: Control who/when doing the work. Using ack-last mechanism.
*/
