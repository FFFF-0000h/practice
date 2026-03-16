package main

import "fmt"

func main() {
	// Test prints
	fmt.Println(union("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
	fmt.Println(union("ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"))
	fmt.Println(union("rien", "cette phrase ne cache rien"))
}

func union(s1, s2 string) string {
	seen := [256]bool{}
	result := ""

	for i := 0; i < len(s1); i++ {
		c := s1[i]
		if !seen[c] {
			seen[c] = true
			result += string(c)
		}
	}

	for i := 0; i < len(s2); i++ {
		c := s2[i]
		if !seen[c] {
			seen[c] = true
			result += string(c)
		}
	}

	return result
}
