package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		v := (int(s[i]) + len(s)) % 127
		if v < 33 {
			v += 33
		}
		r += string(v)
	}
	return r
}
