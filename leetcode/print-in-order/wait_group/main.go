package main

import (
	"fmt"
	"sync"
)

// 0ms Beats 100.00%
type Foo struct {
	wg1, wg2 *sync.WaitGroup
}

func NewFoo() *Foo {
	f := &Foo{
		wg1: &sync.WaitGroup{},
		wg2: &sync.WaitGroup{},
	}
	f.wg1.Add(1)
	f.wg2.Add(1)
	return f
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.wg1.Done()
}

func (f *Foo) Second(printSecond func()) {
	f.wg1.Wait()
	/// Do not change this line
	printSecond()
	f.wg2.Done()
}

func (f *Foo) Third(printThird func()) {
	f.wg2.Wait()
	// Do not change this line
	printThird()
}

func main() {
	f := NewFoo()
	go f.First(func() { fmt.Println("first") })
	go f.Second(func() { fmt.Println("second") })
	f.Third(func() { fmt.Println("third") })
}
