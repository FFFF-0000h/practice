package main

import "github.com/01-edu/z01"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func hex(b byte) {
	h := "0123456789abcdef"
	z01.PrintRune(rune(h[b/16]))
	z01.PrintRune(rune(h[b%16]))
}
func PrintMemory(a [10]byte) {
	for i := 0; i < 10; i++ {
		hex(a[i])
		if i%4 == 3 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, v := range a {
		if v >= 32 && v <= 126 {
			z01.PrintRune(rune(v))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
