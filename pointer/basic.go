package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Basic: & gets address, * dereferences
	x := 42
	p := &x
	*p = 100
	fmt.Println("x =", x)   // 100
	fmt.Println("p =", p)   // address of x
	fmt.Println("*p =", *p) // 100 (dereferenced)

	// Modify via pointer
	num := 5
	increment := func(n *int) {
		*n++
	}
	increment(&num)
	fmt.Println("num =", num) // 6

	// Struct pointer (auto-dereference)
	person := Person{"Alice", 25}
	updatePerson := func(p *Person) {
		p.Age = 30
	}
	updatePerson(&person)
	fmt.Println(person) // {Alice 30}

	// Nil check
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("nil pointer")
	}

	// Return pointer (escape to heap)
	p2 := &Person{"Bob", 35}
	fmt.Println(*p2)
}
