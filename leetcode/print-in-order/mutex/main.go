package main

import (
	"fmt"
	"sync"
)

// 5ms Beats 50.27%
type Foo struct {
	mu1 sync.Mutex
	mu2 sync.Mutex
}

func NewFoo() *Foo {
	instance := &Foo{}
	instance.mu1.Lock()
	instance.mu2.Lock()
	return instance
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.mu1.Unlock()
}

func (f *Foo) Second(printSecond func()) {
	f.mu1.Lock()
	/// Do not change this line
	printSecond()
	f.mu2.Unlock()
}

func (f *Foo) Third(printThird func()) {
	f.mu2.Lock()
	// Do not change this line
	printThird()
}

func main() {
	f := NewFoo()
	go f.First(func() { fmt.Println("first") })
	go f.Second(func() { fmt.Println("second") })
	f.Third(func() { fmt.Println("third") })
}
