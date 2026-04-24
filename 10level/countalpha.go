package main

import "fmt"

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {
	// Convert the string to a rune
	r := []rune(s)
	// Set a counter that ticks for every count
	count := 0

	// Iterate over the slice of runes, neglecting the value
	for i, _ := range r {
		if r[i] >= 'a' && r[i] <= 'z' || r[i] >= 'A' && r[i] <= 'Z' {
			count++
		}
	}
	return count
}
