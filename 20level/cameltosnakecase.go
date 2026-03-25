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

/*
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
*/


func CamelToSnakeCase(s string) string {
	lengthArgs := len(s)

	if lengthArgs == 0 {
		return ""
	}

	// Check for non‑letter characters
	for i := 0; i < lengthArgs; i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= 'A' && s[i] <= 'Z') {
			return s
		}
	}

	// Last character must not be uppercase
	last := s[lengthArgs-1]
	if last >= 'A' && last <= 'Z' {
		return s
	}

	// Check for consecutive uppercase and presence of at least one uppercase
	hasUpperCase := false
	for i := 0; i < lengthArgs-1; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			hasUpperCase = true
		}
		if (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			return s
		}
	}
	if last >= 'A' && last <= 'Z' {
		hasUpperCase = true
	}
	if !hasUpperCase {
		return s
	}

	// Convert to snake_case (preserving case)
	result := ""
	for i, ch := range s {
		if i > 0 && ch >= 'A' && ch <= 'Z' {
			result += "_" + string(ch)
		} else {
			result += string(ch)
		}
	}
	return result
}
