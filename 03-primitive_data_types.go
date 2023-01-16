package main

import "fmt"

func main() {
	// Integer numbers
	// Range: [-2^(bits-1), +2^(bits-1) - 1]
	var a int   // Depends on the OS (32 o 64 bits)
	var b int8  // 8 bits
	var c int16 // 16 bits
	var d int32 // 32 bits
	var e int64 // 64 bits

	fmt.Println(a, b, c, d, e)

	// Positive integer numbers
	// Range: [0, +2^(bits) - 1]
	var f uint   // Depends on the OS (32 o 64 bits)
	var g uint8  // 8 bits
	var h uint16 // 16 bits
	var i uint32 // 32 bits
	var j uint64 // 64 bits

	fmt.Println(f, g, h, i, j)

	// Float numbers
	var k float32 // 32 bits
	var l float64 // 64 bits

	fmt.Println(k, l)

	// Text and boolean
	var m string // use double quotes
	var n bool   // true or false

	fmt.Println(m, n)

	// Complex numbers
	var o complex64  // real and imaginary float32
	var p complex128 // real and imaginary float64

	fmt.Println(o, p)
}
