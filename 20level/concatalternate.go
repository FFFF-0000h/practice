package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(a, b []int) []int {
	res := []int{}
	i := 0
	for i < len(a) || i < len(b) {
		if i < len(a) {
			res = append(res, a[i])
		}
		if i < len(b) {
			res = append(res, b[i])
		}
		i++
	}
	return res
}
