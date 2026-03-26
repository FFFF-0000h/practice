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

/*
package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	// Check if there's exactly one argument
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	
	input := os.Args[1]
	
	// Split by spaces to get words (preserving all non-space characters)
	words := make([]string, 0)
	currentWord := ""
	
	for _, char := range input {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
			// Skip consecutive spaces
		} else {
			currentWord += string(char)
		}
	}
	
	// Add the last word
	if currentWord != "" {
		words = append(words, currentWord)
	}
	
	// If no words found, print newline
	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}
	
	// Print words from index 1 to end
	for i := 1; i < len(words); i++ {
		// Print the word
		for _, char := range words[i] {
			z01.PrintRune(char)
		}
		// Print space between words
		if i < len(words)-1 {
			z01.PrintRune(' ')
		}
	}
	
	// Print the first word at the end
	if len(words) > 1 {
		z01.PrintRune(' ')
	}
	for _, char := range words[0] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
*/

/*
package main

import (
        "fmt"
        "os"
        "strings"
)

func deleteExtraSpaces(a []string) []string {
        var res []string
        for _, v := range a {
                if v != "" {
                        res = append(res, v)
                }
        }
        return res
}

func main() {
        if len(os.Args) == 2 {
                words := strings.Split(os.Args[1], " ")
                words = deleteExtraSpaces(words)
                if len(words) >= 1 {
                        for _, v := range words[1:] {
                                fmt.Print(v, " ")
                        }
                        fmt.Print(words[0])
                }
        }
        fmt.Println()
}
*/
