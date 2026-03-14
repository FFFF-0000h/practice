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

func PrintIfNot(str string) string {
	numstring := len(str)

	if numstring < 3 {
		return "G"
	} else if numstring == 0 {
		return "G\n"
	} else {
		return "Invalid Output\n"
	}
}
