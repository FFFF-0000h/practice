package main

import "fmt"

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord("hello   .........  bye"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("Evidence"))
}

func FirstWord(s string) string {
	// Convert the input string to a slice of runes.
	// Say input is s = "forgive and forget"
	// This ensures we handle Unicode characters correctly (not just ASCII).
	r := []rune(s)

	// word will hold the first word we extract. "forgive"
	word := ""

	// foundWord acts as a flag. It becomes true once we start collecting characters
	// for the first word, so we know when to stop at the next space.
	foundWord := false

	// Iterate over each rune in the slice.
	for i := 0; i < len(r); i++ {
		// If the current rune is a space:
		if r[i] == ' ' {
			// If we have already started collecting a word, this space marks the end.
			if foundWord {
				// Stop the loop entirely – we’ve reached the end of the first word.
				break
			}
			// Otherwise, we haven’t started a word yet, so skip this leading space.
			continue
		}

		// This rune is not a space, so it belongs to the first word.
		// Append it to our word string.
		word += string(r[i])

		// Now that we’ve added a non‑space character, we know a word has begun.
		foundWord = true

		// If we are at the last rune of the slice, there is no following space.
		// The word ends here, so we can return immediately.
		if i == len(r)-1 {
			return word + "\n"
		}
	}

	// If the loop finishes normally (e.g., because we hit a space after the word),
	// we return the collected word plus a newline.
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
