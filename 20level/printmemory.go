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
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for i := 0; i < 10; i++ {
		if a[i] >= 32 && a[i] <= 126 {
			z01.PrintRune(rune(a[i]))
		} else {
			z01.PrintRune('.')
		}
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}
}
