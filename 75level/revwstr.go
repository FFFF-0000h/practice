package main

import "fmt"

func main() {
	// Test prints
	fmt.Println(revwstr("the time of contempt precedes that of indifference"))
	fmt.Println(revwstr("abcdefghijklm"))
	fmt.Println(revwstr("he stared at the mountain"))
	fmt.Println(revwstr(""))
}

func revwstr(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Count words
	words := []string{}
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if start < i {
				words = append(words, s[start:i])
			}
			start = i + 1
		}
	}

	// Build reversed string
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		if i < len(words)-1 {
			result += " "
		}
		result += words[i]
	}

	return result
}


/*
package main

import (
        "fmt"
        "os"
        "strings"
)

func main() {
        if len(os.Args) == 2 {
                a := strings.Split(os.Args[1], " ")
                for i := len(a) - 1; i >= 0; i-- {
                        fmt.Print(a[i])
                        if i != 0 {
                                fmt.Print(" ")
                        }
                }
                fmt.Println()
        }
}
*/
