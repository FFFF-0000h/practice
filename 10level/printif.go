package main

import "fmt"

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(s string) string {
	if len(s) > 3 || s == "" {
		return "G\n"
	}
	return "Invalid Input\n"
}

/*
func PrintIf(str string) string {
	lengthArg := len(str)
	if lengthArg == 0 {
		return "G\n"
	}
	if lengthArg >= 3 {
		return "G\n"
	} else {
		return "Invalid Input"
	}
}
*/
