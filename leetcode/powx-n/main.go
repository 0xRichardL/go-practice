package main

// 0ms Beats 100.00%
// 3.98MB Beats 78.95%
// Non-Recursive, iterative approach using exponentiation by squaring optimization, O(log n)
func myPow_iteration(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	positive := n > 0
	if !positive {
		n = -n
	}
	res := x
	for i := 1; i < n; {
		j := 1
		subRes := x
		for ; j*2 < n-i; j = j * 2 {
			subRes *= subRes
		}
		i += j
		res *= subRes
	}
	if positive {
		return res
	}
	return 1 / res
}
