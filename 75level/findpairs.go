package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPairs("[1, 2, 3, 4, 5]", "6"))
	fmt.Println(findPairs("[-1, 2, -3, 4, -5]", "1"))
	fmt.Println(findPairs("[1, 2, 3, 4, 5]", "10"))
	fmt.Println(findPairs("[-1, -2, -3, -4, -5]", "-5"))
	fmt.Println(findPairs("[1, 2, 3, 4, 20, -4, 5]", "2 5"))
	fmt.Println(findPairs("[1, 2, 3, 4, 20, p, 5]", "5"))
	fmt.Println(findPairs("[1, 2, 3, 4", "5"))
	fmt.Println(findPairs("1, 2, 3, 4", "5"))
}

/*
func findPairs(arr []int, targetSum int) [][]int {
	var pairs [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == targetSum {
				pairs = append(pairs, []int{i, j})
			}
		}
	}
	return pairs
}

func isValidArrayFormat(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 2 || s[0] != '[' || s[len(s)-1] != ']' {
		return false
	}
	return true
} */

func findPairs(a, t string) string {
	if len(a) < 2 || a[0] != '[' || a[len(a)-1] != ']' {
		return "Invalid input."
	}
	n := []int{}
	x, s := 0, 1
	for i := 1; i < len(a)-1; i++ {
		c := a[i]
		if c == ' ' {
			continue
		}
		if c == '-' {
			s = -1
			continue
		}
		if c >= '0' && c <= '9' {
			x = x*10 + int(c-'0')
		} else if c == ',' {
			n = append(n, x*s)
			x, s = 0, 1
		} else {
			return "Invalid number: " + string(c)
		}
	}
	n = append(n, x*s)
	tg, s := 0, 1
	for i := 0; i < len(t); i++ {
		c := t[i]
		if c == ' ' {
			return "Invalid target sum."
		}
		if c == '-' {
			s = -1
		} else if c >= '0' && c <= '9' {
			tg = tg*10 + int(c-'0')
		} else {
			return "Invalid target sum."
		}
	}
	tg *= s
	p := [][2]int{}
	u := make([]bool, len(n))
	for i := 0; i < len(n); i++ {
		if u[i] {
			continue
		}
		for j := i + 1; j < len(n); j++ {
			if !u[j] && n[i]+n[j] == tg {
				p = append(p, [2]int{i, j})
				u[i], u[j] = true, true
				break
			}
		}
	}
	if len(p) == 0 {
		return "No pairs found."
	}
	r := "Pairs with sum "
	if tg < 0 {
		r += "-" + i2a(-tg)
	} else {
		r += i2a(tg)
	}
	r += ": ["
	for i, v := range p {
		if i > 0 {
			r += " "
		}
		r += "[" + i2a(v[0]) + " " + i2a(v[1]) + "]"
	}
	return r + "]"
}

func i2a(n int) string {
	if n == 0 {
		return "0"
	}
	b := []byte{}
	for ; n > 0; n /= 10 {
		b = append([]byte{byte('0' + n%10)}, b...)
	}
	return string(b)
}
