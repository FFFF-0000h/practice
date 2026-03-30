package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for no arguments
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	var options int32 = 0
	
	// Process each argument
	for _, arg := range os.Args[1:] {
		// Check if it's an option (starts with -)
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		
		// Check if the first character after the dash is 'h'
		// This takes priority over all other flags
		if arg[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
		
		// Process each character in the option string
		for i := 1; i < len(arg); i++ {
			c := arg[i]
			
			// Check if character is a valid option (a-z)
			if c < 'a' || c > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			
			// Set the corresponding bit
			// 'a' = bit 0, 'b' = bit 1, ..., 'z' = bit 25
			bitPosition := int(c - 'a')
			options |= 1 << bitPosition
		}
	}
	
	// Print the 32-bit representation in groups of 8 bits
	// From most significant byte (bits 31-24) to least (bits 7-0)
	for i := 3; i >= 0; i-- {
		// Get the current byte
		byte := (options >> (i * 8)) & 0xFF
		
		// Print each bit of the byte from most significant to least
		for j := 7; j >= 0; j-- {
			if (byte>>j)&1 == 1 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		
		// Add space between bytes
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}


/*
package piscine

func Options(args []string) (uint32, bool) {
	var opts uint32 = 0

	for _, arg := range args {
		if len(arg) < 2 || arg[0] != '-' {
			// Not a valid option format
			return 0, false
		}

		for i := 1; i < len(arg); i++ {
			ch := arg[i]
			if ch < 'a' || ch > 'z' {
				// Invalid character
				return 0, false
			}
			// Map 'a'..'z' to bits 0..25
			bit := ch - 'a'
			opts |= 1 << bit
		}
	}

	return opts, true
}

//OR
func Options(args []string) (uint32, bool) {
	var bits uint32 = 0

	for _, arg := range args {
		if len(arg) < 2 || arg[0] != '-' {
			return 0, false
		}

		for i := 1; i < len(arg); i++ {
			c := arg[i]

			if c >= 'a' && c <= 'z' {
				pos := int(c - 'a')
				bits |= 1 << pos
			} else if c == 'h' {
				continue
			} else {
				return 0, false
			}
		}
	}

	return bits, true
}
*/
