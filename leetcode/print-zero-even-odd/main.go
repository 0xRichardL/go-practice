package main

import (
	"fmt"
	"math/rand/v2"
)

type ZeroEvenOdd struct {
	n      int
	v      int
	zeroCh chan struct{}
	evenCh chan struct{}
	oddCh  chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:      n,
		zeroCh: make(chan struct{}, 1),
		evenCh: make(chan struct{}),
		oddCh:  make(chan struct{}),
	}
	zeo.zeroCh <- struct{}{}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for {
		<-z.zeroCh
		if z.v == z.n {
			close(z.zeroCh)
			close(z.evenCh)
			close(z.oddCh)
			return
		}
		printNumber(0)
		z.v++
		if z.v%2 == 0 {
			z.evenCh <- struct{}{}
		} else {
			z.oddCh <- struct{}{}
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for {
		_, ok := <-z.evenCh
		if !ok {
			return
		}
		printNumber(z.v)
		z.zeroCh <- struct{}{}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for {
		_, ok := <-z.oddCh
		if !ok {
			return
		}
		printNumber(z.v)
		z.zeroCh <- struct{}{}
	}
}

func main() {
	zeo := NewZeroEvenOdd(rand.IntN(100))
	printNumber := func(i int) {
		fmt.Print(i)
	}
	go zeo.Even(printNumber)
	go zeo.Odd(printNumber)
	zeo.Zero(printNumber)
}
