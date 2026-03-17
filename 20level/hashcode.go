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

func HashCode(dec string) string {
	size := len(dec)
	hashed := ""
	for _, char := range dec {
		hash := (int(char) + size) % 127
		if hash < 32 || hash > 126 {
			hash += 33
		}
		hashed += string(hash)
	}
	return hashed
}

/*
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
} */
