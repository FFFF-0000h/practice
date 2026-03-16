package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(n int) int {
	for n >= 2 {
		p := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				p = false
				break
			}
		}
		if p {
			return n
		}
		n--
	}
	return 0
}
