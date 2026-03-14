package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

func RetainFirstHalf(str string) string {
	numstring := len(str)

	if numstring == 1 {
		return str
	} else if numstring == 0 {
		return ""
	}
	half := numstring / 2
	return str[:half]
}
