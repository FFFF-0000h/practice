package main

import "fmt"

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	l := len(a)
	start := 0
	end := l

	if len(nbrs) > 0 {
		start = nbrs[0]
		if start < 0 {
			start = l + start
		}
	}

	if len(nbrs) > 1 {
		end = nbrs[1]
		if end < 0 {
			end = l + end
		}
		// If end is 0, it means slice until index 0 (which gives empty slice)
		// But expected output shows nil for invalid ranges
		if end == 0 {
			return nil
		}
	}

	// Check for invalid indices
	if start < 0 || start >= l || end < 0 || end > l || start > end {
		return nil
	}

	return a[start:end]
}
