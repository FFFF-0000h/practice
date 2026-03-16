package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(a []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	res := [][]int{}
	for i := 0; i < len(a); i += size {
		end := i + size
		if end > len(a) {
			end = len(a)
		}
		res = append(res, a[i:end])
	}
	fmt.Println(res)
}
