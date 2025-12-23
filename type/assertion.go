package main

import "fmt"

type StructA struct {
	Value        int
	SpecialValue string
}

func (s *StructA) GetValue() int {
	return s.Value
}

type StructB struct {
	Value int
}

func (s *StructB) GetValue() int {
	return s.Value
}

type Interface interface {
	GetValue() int
}

func basic() {
	var c Interface = &StructA{Value: 10, SpecialValue: "Special"}
	fmt.Println("Value shared by interface: ", c.GetValue())
}

// Unsafe type assertion that may cause panic at runtime
func unsafeTypeAssertion() {
	var c Interface = &StructA{Value: 20, SpecialValue: "Special"}
	// var c Interface = &StructB{Value: 20} // This will cause panic at runtime
	fmt.Println("Special value after unsafe type assertion: ", c.(*StructA).SpecialValue)
}

// Safe type assertion that prevents panic at runtime
func safeTypeAssertion() {
	var c Interface = &StructA{Value: 30, SpecialValue: "Special"}
	if v, ok := c.(*StructA); ok {
		fmt.Println("Special value after safe type assertion: ", v.SpecialValue)
	} else {
		fmt.Println("Type assertion failed")
	}
}

func wrongTypeAssertion() {
	c := &StructB{Value: 40}
	// Uncommenting the following line will cause a compile-time error:
	// "invalid operation: c (variable of type *StructB) is not an interface"
	// c.(*StructA)
}

// Passing interface and determining underlying type
func getTypeFromInterface(i Interface) {
	switch v := i.(type) {
	case *StructA:
		fmt.Println("It's StructA with SpecialValue:", v.SpecialValue)
	case *StructB:
		fmt.Println("It's StructB")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	basic()
	unsafeTypeAssertion()
	safeTypeAssertion()
	getTypeFromInterface(&StructA{})
	getTypeFromInterface(&StructB{})
}
