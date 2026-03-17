package main

import "fmt"

func main() {
	fmt.Println(CountCharacter("Hello World", 'l'))       // 3
	fmt.Println(CountCharacter("5  balloons", 5))         // 1
	fmt.Println(CountCharacter("   ", ' '))               // 3 (if 3 spaces)
	fmt.Println(CountCharacter("The 7 deadly sins", '7')) // 1
}

func CountCharacter(str string, c rune) int {
	counter := 0
	for _, r := range str {
		if r == c {
			counter++
		}
		// When they don't match, do NOTHING - just continue to next character
	}
	return counter
}
