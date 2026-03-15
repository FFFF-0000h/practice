package main

import "github.com/01-edu/z01"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(a [10]byte) {
	h := "0123456789abcdef"

	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(h[a[i]/16]))
		z01.PrintRune(rune(h[a[i]%16]))
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for i := 0; i < 10; i++ {
		c := a[i]
		if c < 32 || c > 126 {
			c = '.'
		}
		z01.PrintRune(rune(c))
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}
}
