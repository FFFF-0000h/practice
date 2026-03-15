package main

import "fmt"

func main() {
	fmt.Println(Itoa(3))
	fmt.Println(Itoa(24))
	fmt.Println(Itoa(599))
}

func Itoa(n int) string {
    // Handle zero
    if n == 0 {
        return "0"
    }
    
    // Handle negative numbers
    isNegative := false
    if n < 0 {
        isNegative = true
        n = -n  // Work with positive number for digit extraction
    }
    
    // Build the string
    result := ""
    for n > 0 {
        // Get the last digit and convert to character
        digit := n % 10
        result = string(rune('0'+digit)) + result
        n = n / 10
    }
    
    // Add minus sign if negative
    if isNegative {
        result = "-" + result
    }
    
    return result
}
