package main

import "fmt"

type Contestant struct {
	Name string
	Age  int
}

// Useful for read-only, non-mutating operations
func (p Contestant) Clone() Contestant {
	return p
}

// Useful for reference, mutating operations
func (p *Contestant) Ref() *Contestant {
	return p
}

func main() {
	p := Contestant{Name: "Alice", Age: 30}
	fmt.Printf("p = %+v\n", p)
	fmt.Println("\nCreate a clone by value receiver")
	clone := p.Clone()
	clone.Name = "Bob"
	fmt.Printf("clone = %+v\n", clone)
	fmt.Println("\nCreate a reference by pointer receiver")
	ref := p.Ref()
	ref.Name = "Charlie"
	fmt.Printf("p = %+v\n", p)
	fmt.Printf("ref = %+v\n", ref)

}
