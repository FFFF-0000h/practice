package main

import "fmt"

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord("hello   .........  bye"))
}

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
}


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
