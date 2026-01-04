package main

import "golang.org/x/exp/constraints"

// constraints.Ordered can be: signed|unsigned integers, floats, or strings
// that supports the operators < <= >= >.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
