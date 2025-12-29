package main

import "fmt"

type Human interface {
	Speak() string
}

type Animal interface {
	Walk() string
}

type Person struct {
}

// Change this to value receiver to see the difference
func (p *Person) Speak() string {
	return "Hello!"
}

// A pointer struct will also have this method in its method set.
func (p Person) Walk() string {
	return "Walking..."
}

func speakOut(h Human) {
	fmt.Println("speakOut:", h.Speak())
}

func walkOut(a Animal) {
	fmt.Println("walkOut:", a.Walk())
}

func main() {
	p := Person{}
	// Still works here, directly calling the methods:
	fmt.Println(p.Walk())
	fmt.Println(p.Speak())
	// Both value and pointer receivers work here (value receiver method)
	walkOut(p)
	walkOut(&p)
	// Only pointer receiver works here
	speakOut(p) // This will cause a compile-time error
	speakOut(&p)
}
