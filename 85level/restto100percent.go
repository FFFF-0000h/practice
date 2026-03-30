//PIGLATIN
package piscine
import "strings"

func PigLatin(str string) string {
	if str == "" {
		return ""
	}

	vowels := "aeiouAEIOU"

	words := strings.Fields(str)
	if len(words) != 1 {
		return ""
	}

	word := words[0]
	firstVowel := -1

	for i, c := range word {
		if strings.ContainsRune(vowels, c) {
			firstVowel = i
			break
		}
	}

	if firstVowel == -1 {
		return "No vowels"
	}

	if firstVowel == 0 {
		return word + "ay"
	}

	return word[firstVowel:] + word[:firstVowel] + "ay"
}

//OR
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	result := pigLatin(args)

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func pigLatin(word string) string {

	//  Check if the word start with a vowel
	if isVowel(word[0]) {
		return word + "ay"
	}

	// Find the position of the first vowel
	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			return word[i:] + word[:i] + "ay"
		}
	}
	return "No Vowels"
}

//ROMANNUMBERS
package piscine

import "github.com/01-edu/z01"

func RomanNumeral(n int) string {
	if n <= 0 || n >= 4000 {
		return "ERROR: cannot convert to roman digit\n"
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	breakdown := ""
	num := n
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			if len(symbols[i]) == 2 {
				breakdown += "(" + string(symbols[i][0]) + "-" + string(symbols[i][1]) + ")+"
			} else {
				breakdown += string(symbols[i]) + "+"
			}
			num -= values[i]
		}
	}
	if len(breakdown) > 0 && breakdown[len(breakdown)-1] == '+' {
		breakdown = breakdown[:len(breakdown)-1]
	}

	roman := ""
	num = n
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	return breakdown + "\n" + roman + "\n"
}

//OR
package main

import (
	"os"
)

var Roman = []struct {
	value   int
	rNumber string
	working string
}{
	{1000, "M", "M"},
	{900, "CM", "(M-C)"},
	{500, "D", "D"},
	{400, "CD", "(D-C)"},
	{100, "C", "C"},
	{90, "XC", "(C-X)"},
	{50, "L", "L"},
	{40, "XL", "(L-X)"},
	{10, "x", "X"},
	{9, "IX", "(X-I)"},
	{5, "V", "V"},
	{4, "IV", "(V-I)"},
	{1, "I", "I"},
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	roman, err := Atoi(os.Args[1])
	if roman <= 0 || roman >= 4000 {
		os.Stdout.WriteString("ERROR: cannot convert to roman digit\n")
		return
	}

	if err != "" {
		os.Stdout.WriteString(err)
		return
	}
	var working string
	var romanOutPut string

	for _, r := range Roman {
		for roman >= r.value {
			if working == "" {
				working += r.working
			} else {
				working += "+" + r.working
			}
			romanOutPut += r.rNumber
			roman -= r.value
		}
	}
	os.Stdout.WriteString(working + "\n" + romanOutPut + "\n")
}

func Atoi(s string) (int, string) {
	var number int
	sign := 1

	for i, v := range s {
		if i == 0 && v == '-' {
			sign = -1
		} else if i == 0 && v == '+' {
			sign = 1
		} else if v >= '0' && v <= '9' {
			number = number*10 + int(v-'0')
		} else {
			return 0, "ERROR: cannot convert to roman digit\n"
		}
	}
	return sign * number, ""
}

//BRACKETS
package piscine

func BracketsCheck(s string) string {
	stack := []rune{}

	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[ch] {
				return "Error\n"
			}
			stack = stack[:len(stack)-1] // pop
		}
	}

	if len(stack) != 0 {
		return "Error\n"
	}
	return "OK\n"
}

//OR
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, args := range os.Args[1:] {
		if checkBrackets(args) {
			printStr("OK")
		} else {
			printStr("Error")
		}
	}
}

func checkBrackets(s string) bool {
	stack := make([]rune, len(s))
	top := -1
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			top++
			stack[top] = c
		} else if c == ')' || c == ']' || c == '}' {
			if top < 0 || stack[top] != pairs[c] {
				return false
			}
			top--
		}
	}
	return top == -1
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

//OR Brackets again
package main

import (
	"fmt"
	"os"
)

func matchBrackets(exp string) bool {
	runes := []rune(exp)
	var opened []rune
	ptr := -1
	for _, c := range runes {
		if c == '(' || c == '[' || c == '{' {
			opened = append(opened, c)
			ptr++
		} else if c == ')' {
			if ptr < 0 || opened[ptr] != '(' {
				return false
			}
			opened = opened[:len(opened)-1]
			ptr--
		} else if c == ']' {
			if ptr < 0 || opened[ptr] != '[' {
				return false
			}
			opened = opened[:len(opened)-1]
			ptr--
		} else if c == '}' {
			if ptr < 0 || opened[ptr] != '{' {
				return false
			}
			opened = opened[:len(opened)-1]
			ptr--
		}
	}
	return len(opened) == 0
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println()
	} else {
		for _, v := range os.Args[1:] {
			if matchBrackets(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("Error")
			}
		}
	}
}




//RPNCALC
package piscine

import "strconv"

func RPN(expr string) string {
	tokens := []string{}
	word := ""
	for _, ch := range expr {
		if ch == ' ' {
			if word != "" {
				tokens = append(tokens, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}
	if word != "" {
		tokens = append(tokens, word)
	}

	if len(tokens) == 0 {
		return "Error\n"
	}

	stack := []int64{}

	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/", "%":
			if len(stack) < 2 {
				return "Error\n"
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int64
			switch tok {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return "Error\n"
				}
				res = a / b
			case "%":
				if b == 0 {
					return "Error\n"
				}
				res = a % b
			}
			stack = append(stack, res)
		default:
			// parse number
			n, err := strconv.ParseInt(tok, 10, 64)
			if err != nil {
				return "Error\n"
			}
			stack = append(stack, n)
		}
	}

	if len(stack) != 1 {
		return "Error\n"
	}

	return strconv.FormatInt(stack[0], 10) + "\n"
}

//OR
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	input := os.Args[1]

	inputSlice := strings.Fields(input)

	var stack []int

	for _, numStr := range inputSlice {
		if num, err := strconv.Atoi(numStr); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				fmt.Println("Error")
                return
			}

			// Pop the last two operands from the stack

			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch numStr {
				case "+":
                    result = a + b
                case "-":
                    result = a - b
                case "*":
                    result = a * b
                case "/":
                    if b == 0 {
                        fmt.Println("Error")
                        return
                    }
                    result = a / b
                default:
                    fmt.Println("Error")
                    return	
			}
			stack = append(stack, result)
		}
	}
	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(stack[0])
}

//BRAINFUCK
package piscine

import (
	"os"
)

func Brainfuck(src string) {
	const memSize = 2048
	mem := [memSize]byte{}
	ptr := 0

	jump := make(map[int]int)
	stack := []int{}
	for i := 0; i < len(src); i++ {
		switch src[i] {
		case '[':
			stack = append(stack, i)
		case ']':
			if len(stack) == 0 {
				return
			}
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			jump[start] = i
			jump[i] = start
		}
	}

	i := 0
	for i < len(src) {
		switch src[i] {
		case '>':
			ptr++
			if ptr >= memSize {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = memSize - 1
			}
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			os.Stdout.Write([]byte{mem[ptr]})
		case '[':
			if mem[ptr] == 0 {
				i = jump[i]
			}
		case ']':
			if mem[ptr] != 0 {
				i = jump[i]
			}
		}
		i++
	}
}

//Brainfuck alt
package main

import (
	"fmt"
	"os"
)

const SIZE = 2048

func main() {
	if len(os.Args) != 2 {
		return
	}
	progpoint := []byte(os.Args[1])
	var arby [SIZE]byte
	pos := 0
	openBr := 0         // opened brackets
	i := 0              // iterates through the source code passed in the argument
	N := len(progpoint) // length of the source code
	for i >= 0 && i < N {
		switch progpoint[i] {
		case '>':
			// Increment the pointer
			pos++
		case '<':
			// decrement the pointes
			pos--
		case '+':
			// increment the pointed byte
			arby[pos]++
		case '-':
			// decrement the pointed byte
			arby[pos]--
		case '.':
			// print the pointed byte on std output
			fmt.Printf("%c", rune(arby[pos]))
		case '[':
			// go to the matching ']' if the pointed byte is 0 (while start)
			openBr = 0
			if arby[pos] == 0 {
				for i < N && (progpoint[i] != byte(']') || openBr > 1) {
					if progpoint[i] == byte('[') {
						openBr++
					} else if progpoint[i] == byte(']') {
						openBr--
					}
					i++
				}
			}
		case ']':
			// go to the matching '[' if the pointed byte is not 0 (while end)
			openBr = 0
			if arby[pos] != 0 {
				for i >= 0 && (progpoint[i] != byte('[') || openBr > 1) {
					if progpoint[i] == byte(']') {
						openBr++
					} else if progpoint[i] == byte('[') {
						openBr--
					}
					i--
				}
			}
		}
		i++
	}
}



//GROUPING
package piscine

import "strings"

func Grouping(pattern, text string) string {
	if len(pattern) == 0 || len(text) == 0 {
		return ""
	}

	if pattern[0] == '(' && pattern[len(pattern)-1] == ')' {
		pattern = pattern[1 : len(pattern)-1]
	}

	subPatterns := strings.Split(pattern, "|")

	words := strings.Fields(text)
	result := ""
	count := 1

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		for _, sp := range subPatterns {
			if sp == "" {
				continue
			}
			if strings.Contains(lowerWord, sp) {
				result += itoa(count) + ": " + word + "\n"
				count++
				break
			}
		}
	}

	return result
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for n > 0 {
		res = string('0'+n%10) + res
		n /= 10
	}
	return res
}

//OR Grouping alt

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func singleSearch(exp []string, text string) []string {
	items := strings.Split(text, " ")
	var result []string

	for _, item := range items {
		for _, word := range exp {
			if strings.Contains(item, word) {
				result = append(result, item)
			}
		}
	}
	return result
}

func simpleSearch(runes []rune, text string) []string {
	exp := string(runes)

	var result []string
	if !strings.ContainsRune(exp, '|') {
		helper := []string{exp}
		result = append(singleSearch(helper, text))
	} else {
		expWords := strings.Split(exp, "|")
		result = append(result, singleSearch(expWords, text)...)
	}
	return result
}

func brackets(regexp, text string) {
	if text == "" || regexp == "" {
		return
	}
	runes := []rune(regexp)

	if runes[0] == '(' && runes[len(runes)-1] == ')' {
		runes = runes[1 : len(runes)-1]
		result := simpleSearch(runes, text)
		for i, s := range result {
			if !unicode.IsLetter(rune(s[len(s)-1])) {
				s = s[0 : len(s)-1]
			}
			fmt.Printf("%d: %s\n", i+1, s)
		}
	}
}

func main() {
	if len(os.Args) == 3 {
		brackets(os.Args[1], os.Args[2])
	}
}
