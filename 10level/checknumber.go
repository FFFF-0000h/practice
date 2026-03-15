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
for _, r := range s {
if r >= '0' && r <= '9' {
return true
}
}
return false
}
