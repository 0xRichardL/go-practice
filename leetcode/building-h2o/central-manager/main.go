package main

type H2O struct {
	registerCh         chan bool     // true - H; false - O
	consHydro, consOxi chan struct{} // Consuming channels
	ackHydro, ackOxi   chan struct{} // Acknowledgement channels
}

func NewH2O() *H2O {
	h := &H2O{
		registerCh: make(chan bool, 1000),
		consHydro:  make(chan struct{}),
		ackHydro:   make(chan struct{}),
		consOxi:    make(chan struct{}),
		ackOxi:     make(chan struct{}),
	}
	go h.manager()
	return h
}

func (h *H2O) manager() {
	hydroThread, oxiThread := 0, 0
	for {
		v := <-h.registerCh
		if v {
			hydroThread++
		} else {
			oxiThread++
		}
		if hydroThread >= 2 && oxiThread >= 1 {
			hydroThread -= 2
			oxiThread--
			h.consHydro <- struct{}{}
			h.consHydro <- struct{}{}
			<-h.ackHydro
			<-h.ackHydro
			h.consOxi <- struct{}{}
			<-h.ackOxi
		}
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.registerCh <- true
	<-h.consHydro
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
	h.ackHydro <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.registerCh <- false
	<-h.consOxi
	// releaseOxygen() outputs "O". Do not change or remove this line.
	releaseOxygen()
	h.ackOxi <- struct{}{}
}
