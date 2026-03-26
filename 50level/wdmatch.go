package main

import "fmt"

func main() {
	// Test prints
	fmt.Println(wdmatch("123", "123"))
	fmt.Println(wdmatch("faya", "fgvvfdxcacpolhyghbreda"))
	fmt.Println(wdmatch("faya", "fgvvfdxcacpolhyghbred"))
	fmt.Println(wdmatch("error", "rrerrrfiiljdfxjyuifrrvcoojh"))
	fmt.Println(wdmatch("quarante deux", "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "))
}

func wdmatch(s1, s2 string) string {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	if i == len(s1) {
		return s1
	}
	return ""
}

/*
package main

import (
    "os"

    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 3 {
        return
    }

    s1 := os.Args[1]
    s2 := os.Args[2]

    if canWrite(s1, s2) {
        printString(s1)
        z01.PrintRune('\n')
    }
}

func canWrite(s1, s2 string) bool {
    i, j := 0, 0

    for i < len(s1) && j < len(s2) {
        if s1[i] == s2[j] {
            i++
        }
        j++
    }

    return i == len(s1)
}

func printString(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
}
*/
