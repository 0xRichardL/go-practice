package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{"Alice", 25}
	num := 42
	pi := 3.14159

	fmt.Printf("%%v\t: %v\n", user)      // {Alice 25}
	fmt.Printf("%%+v\t: %+v\n", user)    // {Name:Alice Age:25}
	fmt.Printf("%%#v\t: %#v\n", user)    // main.User{Name:"Alice", Age:25}
	fmt.Printf("%%T\t: %T\n", user)      // main.User
	fmt.Printf("%%d\t: %d\n", num)       // 42
	fmt.Printf("%%s\t: %s\n", user.Name) // Alice
	fmt.Printf("%%f\t: %f\n", pi)        // 3.141590
	fmt.Printf("%%.2f\t: %.2f\n", pi)    // 3.14
	fmt.Printf("%%.2e\t: %.2e\n", pi)    // 3.14e+00
	fmt.Printf("Width 10: |%10s|%10f|\n", user.Name, pi)
}
