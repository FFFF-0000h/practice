package main

import "fmt"

func main() {
	fmt.Print(LastWord("FOR PONY"))
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord("  "))
	fmt.Print(LastWord("  lorem,ipsum  "))
}

func LastWord(s string) string {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return s[i+1:end+1] + "\n"
}
