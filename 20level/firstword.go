package main

import "fmt"

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord("hello   .........  bye"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("Evidence"))
}

func FirstWord(s string) string {
    r := []rune(s) // Convert the slice to a string
    i := 0
    n := len(r)

    // 1. Skip all leading spaces
    for i < n && r[i] == ' ' {
        i++
    }

    // 2. If we've reached the end, there's no word
    if i == n {
        return "\n"
    }

    // 3. Record the start of the word
    start := i

    // 4. Move forward until we hit a space or the end
    for i < n && r[i] != ' ' {
        i++
    }

    // 5. Extract the word from start to i
    word := string(r[start:i])

    return word + "\n"
}

/*
func FirstWord(s string) string {
	// This converts thestring to a rune and stores it in a variable named 'r'
	r := []rune(s)
	word := ""
	foundWord := false // Switch for if a word has been found

	for i := 0; i < len(r); i++ {
		// check for if the current byte index is an empty space character and also check if a word has been found and if it has, break out of the first if and continue what?
		if r[i] == ' ' {
			if foundWord {
				break
			}
			continue
		}

		// Append word which is an empty string to the the string conversion of the rune byte index
		word += string(r[i])
		foundWord = true

		// check for if i is the last rune, if it is then return word
		if i == len(r)-1 {
			return word + "\n"
		}
	}
	return word + "\n"
}
*/

/*
func FirstWord(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	start := i
	for i < len(s) && s[i] != ' ' {
		i++
	}
	return s[start:i] + "\n"
}
*/

/*
import "strings"

func FirstWord(s string) string {
        words := strings.Fields(s)
        res := "\n"
        if len(words) > 0 {
                res = words[0] + res
        }
        return res
}
*/
