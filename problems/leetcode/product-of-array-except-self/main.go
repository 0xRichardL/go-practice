package main

func productExceptSelf(nums []int) []int {
	zeroCount := 0
	allProduct := 1
	for _, v := range nums {
		if v != 0 {
			allProduct = allProduct * v
		} else {
			zeroCount++
		}
	}
	res := make([]int, 0)
	for _, v := range nums {
		if (v == 0 && zeroCount > 1) || (v != 0 && zeroCount > 0) {
			res = append(res, 0)
			continue
		}
		if v == 0 {
			res = append(res, allProduct)
			continue
		}
		res = append(res, allProduct/v)
	}
	return res
}
