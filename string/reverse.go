package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "abcd"
	fmt.Printf("str: %s\n", str)

	fmt.Printf("reverse(str): %s\n", reverse(str))

	fmt.Printf("str[0:1]: %s\n", str[0:1])
	fmt.Printf("str[0:2]: %s\n", str[0:2])
	fmt.Printf("str[2:2]: %s\n", str[2:2])
	fmt.Printf("str[2:3]: %s\n", str[2:3])
}
