package main

import "fmt"

func main() {
	// Test prints
	fmt.Println(rostring("abc "))
	fmt.Println(rostring("Let there be light"))
	fmt.Println(rostring(" AkjhZ zLKIJz , 23y"))
	fmt.Println(rostring(""))
}

func rostring(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Extract words
	words := []string{}
	start := -1
	for i := 0; i <= len(s); i++ {
		if i < len(s) && s[i] != ' ' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				words = append(words, s[start:i])
				start = -1
			}
		}
	}

	if len(words) == 0 {
		return ""
	}

	// Rotate left: move first word to the end
	if len(words) == 1 {
		return words[0]
	}

	result := words[1]
	for i := 2; i < len(words); i++ {
		result += " " + words[i]
	}
	result += " " + words[0]

	return result
}
