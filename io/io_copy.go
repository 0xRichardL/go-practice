package main

import (
	"io"
	"os"
)

func simpleCopy() {
	// io.File implements both io.Reader and io.Writer interfaces
	// Open a file for reading and writing
	input, err := os.OpenFile("./io/input.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	// copy the content of input file to os.Stdout
	// os.Stdout also implements io.Writer interface
	w, err := io.Copy(os.Stdout, input)
	if err != nil {
		panic(err)
	}
	println("\nBytes written:", w)
}

func bufferedCopy() {
	input, err := os.OpenFile("./io/input.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	// Using io.CopyBuffer to specify a custom buffer
	// 4 bytes at a time
	buf := make([]byte, 4)
	w, err := io.CopyBuffer(os.Stdout, input, buf)
	if err != nil {
		panic(err)
	}
	println("\nBytes written:", w)
}

func apartCopy() {
	input, err := os.OpenFile("./io/input.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	// Copy only first 61 bytes
	io.CopyN(os.Stdout, input, 61)
}

func main() {
	simpleCopy()
	bufferedCopy()
	apartCopy()
}
