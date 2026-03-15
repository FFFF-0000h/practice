package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
    // Handle cases with no possible primes
    if nb < 2 {
        return 0
    }
    
    // Start from nb and go down to find the largest prime
    for i := nb; i >= 2; i-- {
        if isPrime(i) {
            return i
        }
    }
    
    return 0
}

// Optimized prime checking for numbers up to 99999
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    
    // Check odd divisors up to square root
    for i := 3; i*i <= n; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}
