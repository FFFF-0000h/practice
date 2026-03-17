package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}

<<<<<<< HEAD
/* func ConcatSlice(a, b []int) []int {
=======
func ConcatSlice(a, b []int) []int {
      if len(a) == 0 && len(b) == 0 {
        return nil
    }
>>>>>>> 3e7b5064db84b489029c65fee149be743787be62
	return append(a, b...)
} */

func ConcatSlice(slice1, slice2 []int) []int {
	var result []int
	for _, v := range slice1 {
		result = append(result, v)
	}
	for _, v := range slice2 {
		result = append(result, v)
	}
	return result
}
