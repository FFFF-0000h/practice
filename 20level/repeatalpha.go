package main

import "fmt"

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
	fmt.Println(RepeatAlpha(""))
}

func RepeatAlpha(s string) string {
    result := ""
    
    for _, char := range s {
        if char >= 'a' && char <= 'z' {
            // Lowercase letter - repeat based on position
            repeat := int(char - 'a' + 1)
            for i := 0; i < repeat; i++ {
                result += string(char)
            }
        } else if char >= 'A' && char <= 'Z' {
            // Uppercase letter - repeat based on position
            repeat := int(char - 'A' + 1)
            for i := 0; i < repeat; i++ {
                result += string(char)
            }
        } else {
            // Non-letter - add as is
            result += string(char)
        }
    }
    
    return result
}
