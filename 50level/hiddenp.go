package main

import "fmt"

func main() {
	fmt.Println(hiddenp("fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"))
	fmt.Println(hiddenp("abc", "2altrb53c.sse"))
	fmt.Println(hiddenp("abc", "btarc"))
	fmt.Println(hiddenp("DD", "DABC"))
}

func hiddenp(s1, s2 string) int {
	if len(s1) == 0 {
		return 1
	}
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	if i == len(s1) {
		return 1
	}
	return 0
}
