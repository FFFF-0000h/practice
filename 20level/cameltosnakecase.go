package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func containOnlyAlphabet(str string) bool {
	for _, c := range str {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}

func isUpper(s rune) bool {
	return s >= 'A' && s <= 'Z'
}

func CamelToSnakeCase(s string) string {
	result := ""
	if len(s) == 0 || !containOnlyAlphabet(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		if i != 0 && isUpper(rune(s[i])) && i+1 < len(s) && !isUpper(rune(s[i+1])) {
			result += "_"
			result += string(s[i])
		} else if !isUpper(rune(s[i])) || (i == 0 && isUpper(rune(s[i]))) {
			result += string(s[i])
		} else {
			return s
		}
	}
	return result
}

/*
// Helper functions FIRST (outside the main function)
func isLetter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func isUpper(c byte) bool {
	return (c >= 'A' && c <= 'Z')
}

func toLower(c byte) byte {
	return c + 32
}

// Main function AFTER helpers
func CamelToSnakeCase(s string) string {
	numstring := len(s)

	if numstring == 0 {
		return ""
	}

	// Validation pass
	for i := 0; i < numstring; i++ {
		if !isLetter(s[i]) {
			return s
		}

		if i > 0 && isUpper(s[i]) && isUpper(s[i-1]) {
			return s
		}
	}

	if isUpper(s[numstring-1]) {
		return s
	}

	// Conversion pass
	result := ""
	for i := 0; i < len(s); i++ {
		if isUpper(s[i]) {
			if i > 0 {
				result += "_"
			}
			// Keep the original case! Don't lowercase automatically
			result += string(s[i])
		} else {
			result += string(s[i])
		}
	}

	return result
} */
