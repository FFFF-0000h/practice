package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	s := strings.ReplaceAll(str, " ", "")
	var _str strings.Builder
	j := 0
	for _, char := range s {
		if j == 5 {
			_str.WriteRune(rune(' '))
			j = 0
		} else {
			_str.WriteRune(rune(char))
			j++
		}
	}
	_str.WriteRune('\n')
	return _str.String()
}

/*
func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	result, count := "", 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		result += string(str[i])
		count++

		if count == 5 {
			result += " "
			count = 0
			i++ // skip next character (the 6th one)
		}
	}

	// Remove trailing space
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	// Count non-space characters
	validChars := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			validChars++
		}
	}

	if validChars < 5 {
		return "Invalid Input\n"
	}
	return result + "\n"
} */
