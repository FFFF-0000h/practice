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

/*
package main

import (
	"os"
)

func main() {
	// Check if we have exactly 2 arguments
	if len(os.Args) != 3 {
		os.Stdout.Write([]byte("Invalid input.\n"))
		return
	}
	
	arrStr := os.Args[1]
	targetStr := os.Args[2]
	
	// Parse array
	arr, errMsg := parseArray(arrStr)
	if errMsg != "" {
		os.Stdout.Write([]byte(errMsg + "\n"))
		return
	}
	
	// Parse target
	target, errMsg := parseInt(targetStr)
	if errMsg != "" {
		os.Stdout.Write([]byte("Invalid target sum.\n"))
		return
	}
	
	// Find pairs
	pairs := findPairs(arr, target)
	
	// Output result
	if len(pairs) == 0 {
		os.Stdout.Write([]byte("No pairs found.\n"))
		return
	}
	
	output := "Pairs with sum " + itoa(target) + ": ["
	for i, pair := range pairs {
		if i > 0 {
			output += " "
		}
		output += "[" + itoa(pair[0]) + " " + itoa(pair[1]) + "]"
	}
	output += "]\n"
	os.Stdout.Write([]byte(output))
}

func parseArray(s string) ([]int, string) {
	// Check format: must start with [ and end with ]
	if len(s) < 2 || s[0] != '[' || s[len(s)-1] != ']' {
		return nil, "Invalid input."
	}
	
	// Remove brackets
	content := s[1 : len(s)-1]
	if content == "" {
		return []int{}, ""
	}
	
	// Split by commas
	nums := []int{}
	current := ""
	negative := false
	
	for i := 0; i < len(content); i++ {
		c := content[i]
		
		if c == ' ' {
			continue
		}
		
		if c == '-' {
			if current != "" {
				return nil, "Invalid input."
			}
			negative = true
			continue
		}
		
		if c >= '0' && c <= '9' {
			current += string(c)
		} else if c == ',' {
			if current == "" {
				return nil, "Invalid input."
			}
			num := 0
			for j := 0; j < len(current); j++ {
				num = num*10 + int(current[j]-'0')
			}
			if negative {
				num = -num
			}
			nums = append(nums, num)
			current = ""
			negative = false
		} else {
			return nil, "Invalid number: " + string(c)
		}
	}
	
	// Handle last number
	if current != "" {
		num := 0
		for j := 0; j < len(current); j++ {
			num = num*10 + int(current[j]-'0')
		}
		if negative {
			num = -num
		}
		nums = append(nums, num)
	}
	
	return nums, ""
}

func parseInt(s string) (int, string) {
	if s == "" {
		return 0, "Invalid target sum."
	}
	
	negative := false
	start := 0
	
	if s[0] == '-' {
		negative = true
		start = 1
	}
	
	if start >= len(s) {
		return 0, "Invalid target sum."
	}
	
	result := 0
	for i := start; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, "Invalid target sum."
		}
		result = result*10 + int(c-'0')
	}
	
	if negative {
		result = -result
	}
	
	return result, ""
}

func findPairs(nums []int, target int) [][2]int {
	pairs := [][2]int{}
	
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				pairs = append(pairs, [2]int{i, j})
			}
		}
	}
	
	return pairs
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	
	digits := []byte{}
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	
	if negative {
		digits = append([]byte{'-'}, digits...)
	}
	
	return string(digits)
}
*/
