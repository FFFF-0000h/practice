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

/*

import "strings"

func ThirdTimeIsACharm(arg string) string {
        if arg == "" || len(arg) < 3 {
                return "\n"
        }
        var str strings.Builder
        for i := 0; i < len(arg); i++ {
                if i == 0 {
                        continue
                }
                j := i + 1
                if j%3 == 0 {
                        str.WriteRune(rune(arg[i]))
                }                                            }
        str.WriteRune(rune('\n'))
        return (str.String())
}

*/
