package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	r := ""
	c := 0
	for i := 0; i < len(s); i += c {
		c = 0
		for j := i; j < len(s); j++ {
			if s[j] != s[i] {
				break
			}
			c++
		}
		r += Itoa(c) + string(s[i])
	}
	return r
}

func Itoa(n int) (s string) {
	for n > 0 {
		s = string(rune(n%10)+'0') + s
		n /= 10
	}
	return s
}



/*

import "strconv"

func countDuplication(s string, i int) int {
        var count int = 0
        for _, v := range s[i:] {
                if v == rune(s[i]) {
                        count++
                } else {
                        break
                }
        }
        return count
}

func ZipString(s string) string {
        var result string
        i := 0
        for i < len(s) {
                counter := countDuplication(s, i)
                result = result + strconv.Itoa(counter) + string(s[i])
                i += int(counter)
        }
        return result
}

*/
