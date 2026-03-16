package main

import (
	"fmt"
)

func main() {
	// Test cases
	fmt.Println(ItoaBase(10, 2))    // 1010
	fmt.Println(ItoaBase(255, 16))  // FF
	fmt.Println(ItoaBase(42, 8))    // 52
	fmt.Println(ItoaBase(-10, 2))   // -1010
	fmt.Println(ItoaBase(0, 10))    // 0
	fmt.Println(ItoaBase(15, 16))   // F
	fmt.Println(ItoaBase(16, 16))   // 10
	fmt.Println(ItoaBase(123, 4))   // 1323
	fmt.Println(ItoaBase(-255, 16)) // -FF
}

func ItoaBase(v, b int) string {
	base := "0123456789ABCDEF"
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	res := ""
	for v > 0 {
		res = string(base[v%b]) + res
		v /= b
	}
	if neg {
		res = "-" + res
	}
	return res
}
