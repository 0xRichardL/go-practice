package main

import (
	"fmt"
	"os"
)

// NOTE: Build & run with cmd:
// go build -o cli/basic cli/basic.go
// ./cli/basic arg1 arg2 arg3
func main() {
	args := os.Args
	// Expected:
	// args[0] = program path
	// args[1:] = command line arguments
	fmt.Printf(`Args: %#v\n`, args)
}
