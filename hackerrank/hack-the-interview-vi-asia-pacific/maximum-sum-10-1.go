package main

import "fmt"

func performOperations(N int32, op []int32) []int64 {
	bwSum := int64(N-1+2) * int64(N-1-2+1) / 2
	l, r := int32(1), N
	res := []int64{}
	for _, p := range op {
		if p == l || p == r || (p >= 2 && p <= N-1) {
			// in case
			l, r = r, l
		} else {
			// out case
			r = p
		}
		res = append(res, int64(l)+int64(bwSum)+int64(r))
	}
	return res
}

func main() {
	type Test struct {
		N  int32
		Op []int32
	}
	tests := []Test{
		{3, []int32{4, 2}},
	}
	for _, test := range tests {
		output := performOperations(test.N, test.Op)
		fmt.Println(output)
	}
}
