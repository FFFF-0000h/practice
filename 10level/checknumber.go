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
	r := []rune(s) // Convert the string to a slice of runes
	for i := 0; i < len(r); i++ {
		// Iterate over the slice of runes
		if r[i] >= '0' && r[i] <= '9' {
			return true
		}
	}
	return false
}

/*
func main() {
	str := "Hello"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if r[i] >= '0' && r[i] <= '9' {
			fmt.Println(0)
		}
	}
	fmt.Println(1)
}
*/

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
