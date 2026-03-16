package main

import "os"

func main() {
	printrevcomb()
}

func printrevcomb() {
	first := true
	result := ""
	for a := 9; a >= 0; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				if a > b && b > c {
					if !first {
						result += ", "
					}
					result += string(rune('0'+a)) + string(rune('0'+b)) + string(rune('0'+c))
					first = false
				}
			}
		}
	}
	os.Stdout.WriteString(result + "\n")
}

/* package main

import "fmt"

func main() {
	printrevcomb()
}

func printrevcomb() {
	first := true
	for a := 9; a >= 0; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				if a > b && b > c {
					if !first {
						fmt.Print(", ")
					}
					fmt.Printf("%d%d%d", a, b, c)
					first = false
				}
			}
		}
	}
	fmt.Println()
} */

/* package main

func main() {
	printrevcomb()
}

func printrevcomb() {
	first := true
	for a := 9; a >= 0; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				if a > b && b > c {
					if !first {
						print(", ")
					}
					print(string(rune('0'+a)) + string(rune('0'+b)) + string(rune('0'+c)))
					first = false
				}
			}
		}
	}
	print("\n") // Explicit newline
} */

/* package main

func main() {
	printrevcomb()
}

func printrevcomb() {
	first := true
	for a := 9; a >= 0; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				if a > b && b > c {
					if !first {
						print(", ")
					}
					print(string(rune('0'+a)) + string(rune('0'+b)) + string(rune('0'+c)))
					first = false
				}
			}
		}
	}
	println()
} */

/* package main

func main() {
	printrevcomb()
}

func printrevcomb() {
	first := true
	for a := 9; a >= 0; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				if a > b && b > c {
					if !first {
						print(", ")
					}
					print(string(rune('0'+a)) + string(rune('0'+b)) + string(rune('0'+c)))
					first = false
				}
			}
		}
	}
	println() // This should work with cat -e
} */
