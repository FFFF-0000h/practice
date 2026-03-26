package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	// Print hex values in 4-byte rows
	for i := 0; i < 10; i++ {
		b := arr[i]
		fmt.Printf("%c%c", hex[b>>4], hex[b&0x0f])

		if i == 3 || i == 7 {
			fmt.Println()
		} else if i < 9 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	// Print ASCII representation
	for i := 0; i < 10; i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			fmt.Printf("%c", arr[i])
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

/*
package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	r := []string{}
	for _, byt := range arr {
		if byt == 0 {
			r = append(r, "00")
		} else {
			r = append(r, hex(int(byt)))
		}
	}

	// printing r:
	for i := 0; i < len(r); i++ {
		ppstr(r[i])
		// Only print space if:
		// 1. It's not the last element in the array
		// 2. It's not the last element in a row (position 3, 7)
		if i < len(r)-1 && (i+1)%4 != 0 {
			z01.PrintRune(' ')
		}
		if (i+1)%4 == 0 || i == len(r)-1 {
			z01.PrintRune('\n')
		}
	}

	// printing ASCII graphic characters
	for _, byt := range arr {
		if byt < 32 || byt > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(byt))
		}
	}
	z01.PrintRune('\n')
}

func ppstr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func hex(n int) string {
	if n == 0 {
		return "00"
	}
	base := "0123456789abcdef"
	sliceBase := []rune(base)
	index := []int{}
	var r []rune
	for n > 0 {
		index = append(index, n%16)
		n /= 16
	}
	// Ensure 2 digits for values < 16
	if len(index) == 1 {
		r = append(r, '0')
	}
	for i := len(index) - 1; i >= 0; i-- {
		r = append(r, sliceBase[index[i]])
	}
	return string(r)
}
*/
