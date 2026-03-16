package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(a, b string) int {
	if a == "" && b == "" {
		return -1
	}
	m := make(map[rune]int)
	for _, r := range a {
		m[r] |= 1 // mark as present in first string
	}
	for _, r := range b {
		m[r] |= 2 // mark as present in second string
	}
	count := 0
	for _, v := range m {
		if v == 1 || v == 2 { // present in exactly one string
			count++
		}
	}
	return count
}
