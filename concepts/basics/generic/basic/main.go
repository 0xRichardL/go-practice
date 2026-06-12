package main

// Generic function that accepts any type and returns it
func accept_any_type[T any](a T) T {
	return a
}

// Generic function that accepts only integer types and increases the value by 1
func integer_increase[T int | int8 | int16 | int32 | int64](a T) T {
	return a + 1
}

// Custom constraint
type Number interface {
	int | float64
}

func double_value[T Number](a T) T {
	return a * 2
}

func max_in_array[T Number](arr []T) (max T, ok bool) {
	if len(arr) == 0 {
		// Both works:
		// Construct a zero value of type T
		var zero T
		return zero, false
		// Auto zero value by inference
		return 0, false
	}
	max = arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max, true
}

// comparable means types that support == and != operators
func first_match[T comparable](arr []T, target T) (index int, ok bool) {
	for i, v := range arr {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

func main() {
	// Demo accept_any_type
	println("accept_any_type:")
	println(accept_any_type(42))
	println(accept_any_type("hello"))
	println(accept_any_type(3.14))

	// Demo integer_increase
	println("\ninteger_increase:")
	println(integer_increase(10))
	println(integer_increase(int32(100)))

	// Demo double_value
	println("\ndouble_value:")
	println(double_value(5))
	println(double_value(2.5))

	// Demo max_in_array
	println("\nmax_in_array:")
	intArr := []int{3, 7, 2, 9, 1}
	maxInt, ok := max_in_array(intArr)
	println("Max:", maxInt, "Found:", ok)

	floatArr := []float64{1.5, 3.7, 2.1, 9.9, 0.5}
	maxFloat, ok := max_in_array(floatArr)
	println("Max:", maxFloat, "Found:", ok)

	emptyArr := []int{}
	_, ok = max_in_array(emptyArr)
	println("Empty array, Found:", ok)

	// Demo first_match
	println("\nfirst_match:")
	strArr := []string{"apple", "banana", "cherry", "banana"}
	idx, ok := first_match(strArr, "banana")
	println("Index:", idx, "Found:", ok)

	idx, ok = first_match(intArr, 9)
	println("Index:", idx, "Found:", ok)

	idx, ok = first_match(intArr, 99)
	println("Index:", idx, "Found:", ok)
}
