package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	input, err := os.OpenFile("./io/input.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	io.Pipe()
	// Remove the duplicate io.Pipe() call above
	// The piping logic is already implemented below:
	// 1. io.Pipe() creates a connected reader/writer pair
	// 2. The goroutine writes data from input file to the pipe's writer
	// 3. The main goroutine reads from the pipe's reader and outputs to stdout
	r, w := io.Pipe()

	// Write to the pipe in a goroutine
	go func() {
		defer w.Close()
		for i := 2; i > 0; i-- {
			fmt.Println("Count down to start:", i)
			time.Sleep(1 * time.Second)
		}
		io.Copy(w, input)
	}()

	// Read from the pipe and copy to stdout
	io.Copy(os.Stdout, r)
}
