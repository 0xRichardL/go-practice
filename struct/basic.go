package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Struct literals: Declaration & initialization
	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{"Bob", 30} // positional
	var p3 Person           // zero values: "" and 0
	p4 := &Person{"Carol", 28}

	fmt.Println(p1, p2, p3, p4)

	// Field access & modification
	p1.Age = 26
	fmt.Println("Modified:", p1)

	// Struct literals: Anonymous struct (one-off use)
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Println("Config:", config)

	// Struct literals (shorthand init)
	users := []Person{
		{"Dan", 22},
		{"Eve", 24},
	}
	fmt.Println("Users:", users)

	// Comparing structs (must be comparable types)
	a := Person{"Tom", 20}
	b := Person{"Tom", 20}
	c := Person{"Tom", 21}
	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false
	// Note: structs with slices/maps are not comparable
}
