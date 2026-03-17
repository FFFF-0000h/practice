package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	newWord := true
	for _, r := range s {
		if newWord {
			if r >= 'a' && r <= 'z' {
				return false
			}
			newWord = false
		}
		if r == ' ' {
			newWord = true
		}
	}
	return true
}

/*

func IsCapitalized(s string) bool {
        if len(s) == 0 {
                return false
        }

        for i := 0; i < len(s); i++ {
                if s[i] >= 'a' && s[i] <= 'z' && i != 0 && s[i-1] == ' ' {
                        return false
                }
        }
        return !(s[0] >= 'a' && s[0] <= 'z')
}

*/
