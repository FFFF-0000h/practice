package main

import "fmt"

func main() {
	// Test prints
	fmt.Println(union("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
	fmt.Println(union("ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"))
	fmt.Println(union("rien", "cette phrase ne cache rien"))
}

func union(s1, s2 string) string {
	seen := [256]bool{}
	result := ""

	for i := 0; i < len(s1); i++ {
		c := s1[i]
		if !seen[c] {
			seen[c] = true
			result += string(c)
		}
	}

	for i := 0; i < len(s2); i++ {
		c := s2[i]
		if !seen[c] {
			seen[c] = true
			result += string(c)
		}
	}

	return result
}

/*
THIS WORKED FOR UNION:
package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	
	// Check if we have exactly 2 arguments
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}
	
	// Get both strings
	str1 := args[0]
	str2 := args[1]
	
	// Track seen characters
	seen := make(map[rune]bool)
	
	// Process first string - maintain order
	for _, char := range str1 {
		if !seen[char] {
			seen[char] = true
			z01.PrintRune(char)
		}
	}
	
	// Process second string - only add if not seen
	for _, char := range str2 {
		if !seen[char] {
			seen[char] = true
			z01.PrintRune(char)
		}
	}
	
	z01.PrintRune('\n')
}
*/
