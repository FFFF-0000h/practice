package main

import "fmt"

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(s string) string {
	res := ""
	for i := 2; i < len(s); i += 3 {
		res += string(s[i])
	}
	return res + "\n"
}
