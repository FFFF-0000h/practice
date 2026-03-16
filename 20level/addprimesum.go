package main

import "fmt"

func main() {
	fmt.Println(AddPrimeSum(5))
	fmt.Println(AddPrimeSum(7))
	fmt.Println(AddPrimeSum(-2))
	fmt.Println(AddPrimeSum(0))
}

func AddPrimeSum(n int) int {
	if n <= 0 {
		return 0
	}
	s := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			s += i
		}
	}
	return s
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}
