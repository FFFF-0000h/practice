package main

import (
	"fmt"
)

func main() {
	// Test prints
	fmt.Println(reversestrcap("First SMALL TesT"))
	fmt.Println(reversestrcap("SEconD Test IS a LItTLE EasIEr"))
	fmt.Println(reversestrcap("bEwaRe IT'S NoT HARd WhEN "))
	fmt.Println(reversestrcap(" Go a dernier 0123456789 for the road e"))
}

func reversestrcap(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 'a' - 'A'
		}
	}
	for i := 0; i < len(b); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			for i < len(b) && ((b[i] >= 'a' && b[i] <= 'z') || (b[i] >= 'A' && b[i] <= 'Z')) {
				i++
			}
			if i-1 >= 0 && b[i-1] >= 'a' && b[i-1] <= 'z' {
				b[i-1] -= 'a' - 'A'
			}
		}
	}
	return string(b)
}
