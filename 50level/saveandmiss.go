package main

import (
	"fmt"
)

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}

func SaveAndMiss(arg string, num int) string {
	if num <= 0 || num > len(arg) {
		return string(arg)
	}
	_str := ""
	for i := 0; i < len(arg); i++ {
		if i != 0 && i%num == 0 {
			i += num
			if i > len(arg)-1 {
				break
			}
		}
		if i != len(arg) {
			_str += string(rune(arg[i]))
		}
	}
	return _str
}

/*
func SaveAndMiss(s string, n int) string {
	if n <= 0 {
		return s
	}
	res := ""
	save := true
	count := 0
	for _, r := range s {
		if save {
			res += string(r)
		}
		count++
		if count == n {
			save = !save
			count = 0
		}
	}
	return res
} */
