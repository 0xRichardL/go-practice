package main

import (
	"fmt"
	"strconv"
)

// string conversions and parsing
func main() {
	fmt.Println("=== String Conversions & Parsing ===")

	// String to int
	i, _ := strconv.Atoi("42")
	fmt.Printf("Atoi: %d\n", i)

	// String to int64 with base
	i64, _ := strconv.ParseInt("101010", 2, 64) // binary
	fmt.Printf("ParseInt (base 2): %d\n", i64)

	// Int to string
	s := strconv.Itoa(123)
	fmt.Printf("Itoa: %s\n", s)

	// Int64 to string with base
	hex := strconv.FormatInt(255, 16)
	fmt.Printf("FormatInt (hex): %s\n", hex)

	// String to float
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Printf("ParseFloat: %f\n", f)

	// Bool conversions
	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool: %t\n", b)
	fmt.Println()
}
