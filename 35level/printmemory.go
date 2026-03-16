/* package main

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	// Print hex values in 4-byte rows
	for i := 0; i < 10; i++ {
		b := arr[i]
		print(string(hex[b>>4]) + string(hex[b&0x0f]))

		if i == 3 || i == 7 {
			print("\n")
		} else if i < 9 {
			print(" ")
		}
	}
	print("\n")

	// Print ASCII representation
	for i := 0; i < 10; i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			print(string(arr[i]))
		} else {
			print(".")
		}
	}
	print("\n")
} */

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
