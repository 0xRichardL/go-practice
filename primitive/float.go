package main

import (
	"fmt"
	"math"
)

// Key difference: .5 rounding behavior
// Round: .5 away from zero
// RoundToEven: .5 to nearest even (banker's rounding)
func rounding() {
	tests := []float64{2.5, 3.5, -2.5}
	for _, x := range tests {
		fmt.Printf("%.1f: Round=%.0f, RoundToEven=%.0f\n",
			x, math.Round(x), math.RoundToEven(x))
	}
}

func main() {
	rounding()
}
