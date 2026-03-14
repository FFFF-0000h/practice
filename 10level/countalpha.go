package main

import "fmt"

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {
	numstring := len(s)
	counter := 0
	for i := 0; i < numstring; i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {
			counter++
		}
	}
	return counter
}
