package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}

/*
func PrintIfNot(s string) string {
	if len(s) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
*/

func PrintIfNot(str string) string {
	lengthArgs := len(str)
	if lengthArgs < 3 {
		return "G\n"
	}
	if lengthArgs == 0 {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}
