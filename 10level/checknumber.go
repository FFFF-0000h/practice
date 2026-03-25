package main

import "fmt"

func main() {
	fmt.Println(CheckNum("Hello"))           // false
	fmt.Println(CheckNum("Hello1"))          // true
	fmt.Println(CheckNum("123"))             // true
	fmt.Println(CheckNum(""))                // false (empty string)
	fmt.Println(CheckNum("Hello World"))     // false
	fmt.Println(CheckNum("Hello 123 World")) // true
}

func CheckNum(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] >= '0' && r[i] <= '9' {
			return true
		}
	}
	return false
}

/*
func CheckNumber(arg string) bool {
	r := []rune(arg)
	for i, _ := range r {
		if r[i] >= '0' && r[i] <= '9' {
			return true
		}
	}
	return false
}
*/
