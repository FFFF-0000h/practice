package main

import "fmt"

func main() {
	fmt.Println(FirstWord("hello there"))
	fmt.Println(FirstWord("hello   .........  bye"))
	fmt.Println(FirstWord("hello\" \"there"))
}

func FirstWord(sit string) string {
	numstring := len(sit)

	for i := 0; i < numstring; i++ {
		if sit[i] == ' ' {
			return sit[:i]
		}
	}
	// If no space found, return the whole string
	return sit
}
