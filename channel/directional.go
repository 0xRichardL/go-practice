package main

// Type-safe, only allow sending.
func sendOnly(ch chan<- int, value int) {
	ch <- value
}

// Type-safe, only allow receiving.
func receiveOnly(ch <-chan int) int {
	return <-ch
}

func main() {
	ch := make(chan int)

	go sendOnly(ch, 42)
	value := receiveOnly(ch)
	println("Received value:", value)
}
