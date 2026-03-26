package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(arg string) string {
	if arg == "" {
		return "Invalid Output\n"
	}
	var str []string = strings.Split(arg, " ")
	_len := len(str)
	var str1 string
	for i := _len - 1; i >= 0; i-- {
		if len(str[i]) != 0 {
			str1 += str[i]
		}
		if i > 0 && len(str[i-1]) != 0 {
			str1 += " "
		}
	}
	return (strings.TrimSpace(str1) + "\n")
}

/*
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
} */

/*
package solutions

import "strings"

func WordFlip(arg string) string {
	if arg == "" {
		return "Invalid Output\n"
	}
	
	// Split by spaces, preserving empty strings from consecutive spaces
	words := strings.Split(arg, " ")
	
	// Build result with words in reverse order
	var result strings.Builder
	
	// Find first non-empty word from the end to start building
	firstWord := true
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			if !firstWord {
				result.WriteString(" ")
			}
			result.WriteString(words[i])
			firstWord = false
		}
	}
	
	// If all words were empty (only spaces), return newline
	if result.Len() == 0 {
		return "\n"
	}
	
	return result.String() + "\n"
}
*/
