package main

import "fmt"

func main() {
	str := "123456"
	// Direct access to each character in the string
	// Result is in byte format
	for i := 0; i < len(str); i++ {
		var char byte = str[i]
		fmt.Println(char)
	}

	// Using range to iterate over each character in the string
	// Result is in rune format
	for _, char := range str {
		fmt.Println(char)
	}
}
