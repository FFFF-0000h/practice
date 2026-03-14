package main

import "fmt"

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(str string) string {
	numstring := len(str)
	if numstring > 3 {
		return "G\n"
	} else if numstring == 0 {
		return "G\n"
	} else {
		return "Invalid Output\n"
	}
}
