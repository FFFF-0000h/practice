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

func ifNegative(a []string, n int) int {
	if n < 0 {
		n = len(a) + n
	}

	if n < 0 {
		n = 0
	} else if n > len(a) {
		n = len(a)
	}

	return n
}

func Slice(a []string, nbr ...int) []string {
	if len(nbr) == 0 {
		return a
	}

	first := nbr[0]
	if len(nbr) == 1 {
		if first < 0 {
			first = len(a) + first
			if first < 0 {
				return a
			}
		}
		return a[first:]
	}
	second := nbr[1]

	first = ifNegative(a, first)
	second = ifNegative(a, second)

	if first > second {
		return nil
	}

	return a[first:second]
}

/*
func Slice(a []string, nbrs ...int) []string {
	var start, end = 0, len(a)

	if len(nbrs) > 0 {
		start = clampIndex(nbrs[0], len(a))
	}

	if len(nbrs) > 1 {
		end = clampIndex(nbrs[1], len(a))
	}

	if start > end {
		return nil
	}

	return a[start:end]
}

func clampIndex(idx, length int) int {
	if idx < 0 {
		idx += length
	}
	return max(0, min(idx, length))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

/*
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
} */
