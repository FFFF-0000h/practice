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

/*
func ItoaBase(value, base int) string {
	digits := "0123456789ABCDEF"

	if value == 0 {
		return "0"
	}

	neg := value < 0
	if neg {
		value = -value
	}

	res := ""
	for value > 0 {
		res = string(digits[value%base]) + res
		value /= base
	}

	if neg {
		res = "-" + res
	}

	return res
}

//OR
func ItoaBase(value, base int) string {
	digits := "0123456789ABCDEF"

	if value == 0 {
		return "0"
	}

	neg := value < 0
	if neg {
		value = -value
	}

	var buf []byte
	for value > 0 {
		buf = append(buf, digits[value%base])
		value /= base
	}

	// reverse buf
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	if neg {
		return "-" + string(buf)
	}
	return string(buf)
}
*/

/*
func main() {
	baseChars := "0123456789ABCDEF"

	tests := []struct {
		value int
		base  int
	}{
		{10, 2},
		{255, 16},
		{42, 8},
		{-10, 2},
		{0, 10},
		{15, 16},
		{16, 16},
		{123, 4},
		{-255, 16},
	}

	for _, t := range tests {
		v := t.value
		b := t.base

		if v == 0 {
			fmt.Println("0")
			continue
		}

		neg := v < 0
		if neg {
			v = -v
		}

		res := ""
		for v > 0 {
			res = string(baseChars[v%b]) + res
			v /= b
		}

		if neg {
			res = "-" + res
		}

		fmt.Println(res)
	}
}
*/
