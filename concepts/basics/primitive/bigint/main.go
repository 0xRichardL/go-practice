package main

import "math/big"

func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("123456789012345678901234567890", 10)
	b.SetString("987654321098765432109876543210", 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(b, a)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Int).Div(b, a)

	println("Sum: " + sum.String())
	println("Difference: " + diff.String())
	println("Product: " + prod.String())
	println("Quotient: " + quot.String())
}
