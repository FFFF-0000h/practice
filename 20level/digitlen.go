package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

func DigitLen(n, base int) int {
	// base must be between EITHER 2 OR 36, if either of them is true, return -1
	if base < 2 || base > 36 {
		return -1
	}
	
	// if the first integer is less than 0; negative. Negate the negation; turn positive
	if n < 0 {
		n = -n
	}

	count := 0 //counter

	// Core of the program, if the first integer is greater than 0, keep on dividing and counting each time it does it.
	// For example: 100 and 10.
	// 100 divides by 10 3 times until the value turns 0
	// 100 / 10 = 10 ; 1 count
	// 10 (the quotient of the 100 / 10 division)
	// 10 / 10 = 1; 2 counts
	// 2 (the quotient of the 10 / 10 division)
	// 2 / 10 = 0 (with a remainder because 2 cannot be divided by 10) ; 3 counts

	// Therefore, the times or the digitlen (or digitlength) will be 3 for this, validating the algorithm for all other examples that may be fed to the program
	for n > 0 {
		n /= base
		count++
	}
	return count
}
