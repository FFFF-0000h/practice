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
    // Track if we're at the start of a word
    atWordStart := true
    hasAnyWord := false

    // Convert string to runes to handle Unicode characters properly
    runes := []rune(s)

    for i := 0; i < len(runes); i++ {
        char := runes[i]

        if char == ' ' {
            atWordStart = true
            continue
        }

        // We found a non-space character
        hasAnyWord = true

        if atWordStart {
            // Check if it's a lowercase letter (a-z)
            if char >= 'a' && char <= 'z' {
                return false
            }
            atWordStart = false
        }
    }

    return hasAnyWord
}
