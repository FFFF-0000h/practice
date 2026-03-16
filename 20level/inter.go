package main

import "fmt"

func main() {
	// Test prints only
	fmt.Println(inter("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
	fmt.Println(inter("ddf6vewg64f", "twthgdwthdwfteewhrtag6h4ffdhsd"))
}

func inter(s1, s2 string) string {
	seen := [256]bool{}
	for i := 0; i < len(s2); i++ {
		seen[s2[i]] = true
	}

	added := [256]bool{}
	result := ""

	for i := 0; i < len(s1); i++ {
		c := s1[i]
		if seen[c] && !added[c] {
			result += string(c)
			added[c] = true
		}
	}

	return result
}
