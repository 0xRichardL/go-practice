package main

import (
	"fmt"
)

func main() {
	/// Integer types
	// int type is a platform-dependent type (32 or 64 bits)
	var i int = 42
	fmt.Println("int:", i)
	var i8 int8 = -128
	fmt.Println("int8:", i8)
	var i16 int16 = 32767
	fmt.Println("int16:", i16)
	var i32 int32 = 2147483647
	fmt.Println("int32:", i32)
	var i64 int64 = 9223372036854775807
	fmt.Println("int64:", i64)

	/// Unsigned integer types
	// uint type is a platform-dependent type (32 or 64 bits)
	var u uint = 42
	fmt.Println("uint:", u)
	var u8 uint8 = 255
	fmt.Println("uint8:", u8)
	var u16 uint16 = 65535
	fmt.Println("uint16:", u16)
	var u32 uint32 = 4294967295
	fmt.Println("uint32:", u32)
	var u64 uint64 = 18446744073709551615
	fmt.Println("uint64:", u64)

	/// Floating-point types
	var f32 float32 = 3.14
	fmt.Println("float32:", f32)
	var f64 float64 = 3.141592653589793
	fmt.Println("float64:", f64)

	/// Boolean type
	var b bool = true
	fmt.Println("bool:", b)
	// String type
	var s string = "Hello, Go!"
	fmt.Println("string:", s)
}
