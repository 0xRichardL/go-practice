package main

import "os"

func main() {
	input, err := os.OpenFile("./io/input.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	// TODO: Read from file => pipe to change '/' to '\' => write to os.Stdout
}
