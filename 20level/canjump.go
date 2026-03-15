package main

import (
	"fmt"
)

func main() {
	fmt.Println(CanJump([]uint{2, 3, 1, 1, 4}))
	fmt.Println(CanJump([]uint{3, 2, 1, 0, 4}))
	fmt.Println(CanJump([]uint{0}))
}

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}
	MaxReach := 0
	for i, n := range steps {
		if n != 0 {
			MaxReach = i + int(n)
		}
		if MaxReach == len(steps)-1 || len(steps) == 1 {
			return true
		}
	}
	return false
}
