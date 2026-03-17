package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)
	maxLen := len1
	result := make([]int, 0, len1+len2)

	if len2 > maxLen {
		maxLen = len2
	}

	for i := maxLen; i >= 0; i-- {
		if i < len1 {
			result = append(result, slice1[i])
		}
		if i < len2 {
			result = append(result, slice2[i])
		}
	}

	return result
}

/*
func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}
	len1, len2 := len(slice1), len(slice2)

	if len1 > len2 {
		// Add extra elements from larger slice (slice1)
		for i := len1 - 1; i >= len2; i-- {
			result = append(result, slice1[i])
		}
		// Alternate remaining elements
		for i := len2 - 1; i >= 0; i-- {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	} else if len2 > len1 {
		// Add extra elements from larger slice (slice2)
		for i := len2 - 1; i >= len1; i-- {
			result = append(result, slice2[i])
		}
		// Alternate remaining elements
		for i := len1 - 1; i >= 0; i-- {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	} else {
		// Equal length - alternate starting with first slice
		for i := len1 - 1; i >= 0; i-- {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	}

	return result
} */
