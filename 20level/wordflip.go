package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(s string) string {
	if s == "" {
		return "Invalid Output\n" // Added newline here
	}
	words := []string{}
	w := ""
	for _, r := range s {
		if r == ' ' {
			if w != "" {
				words = append(words, w)
				w = ""
			}
		} else {
			w += string(r)
		}
	}
	if w != "" {
		words = append(words, w)
	}
	if len(words) == 0 {
		return "\n"
	}
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		if i > 0 {
			res += " "
		}
	}
	return res + "\n"
}
