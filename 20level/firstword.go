package main

import "fmt"

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord("hello   .........  bye"))
	fmt.Print(FirstWord(""))
}

func FirstWord(s string) string {
	r := []rune(s)
	word := ""
	foundWord := false

	for i := 0; i < len(r); i++ {
		if r[i] == ' ' {
			if foundWord {
				break
			}
			continue
		}
		word += string(r[i])
		foundWord = true
		if i == len(r) - 1 {
			return word + "\n"
		}
	}
	return word + "\n"
}

/*
func FirstWord(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	start := i
	for i < len(s) && s[i] != ' ' {
		i++
	}
	return s[start:i] + "\n"
} */

/*

import "strings"

func FirstWord(s string) string {
        words := strings.Fields(s)
        res := "\n"
        if len(words) > 0 {
                res = words[0] + res
        }
        return res
}

*/
