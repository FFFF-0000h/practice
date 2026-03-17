package main

import "fmt"

func main() {
	fmt.Println(Gcd(42, 26))
	fmt.Println(Gcd(43, 12))
	fmt.Println(Gcd(42, 13))
	fmt.Println(Gcd(14, 6))
	fmt.Println(Gcd(17, 4))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b uint) uint {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

/*
func Gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
} */
