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
	lengthArgs := len(str)
	if lengthArgs == 0 {
		return ""
	}
	if lengthArgs == 1 {
		return str
	}
	return str[:lengthArgs/2]
}

/*

import "strings"

func RetainFirstHalf(str string) string {
        if len(str) == 0 {
                return ""
        } else if len(str) == 1 {
                return (str)
        } else {
                var res strings.Builder
                i := 0
                for i = 0; i < int(len(str)/2); i++ {
                        res.WriteRune(rune(str[i]))
                }
                return res.String()
        }
}

*/
