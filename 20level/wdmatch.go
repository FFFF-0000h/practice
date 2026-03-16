package main

import "fmt"

func main() {
	// Test prints
	fmt.Println(wdmatch("123", "123"))
	fmt.Println(wdmatch("faya", "fgvvfdxcacpolhyghbreda"))
	fmt.Println(wdmatch("faya", "fgvvfdxcacpolhyghbred"))
	fmt.Println(wdmatch("error", "rrerrrfiiljdfxjyuifrrvcoojh"))
	fmt.Println(wdmatch("quarante deux", "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "))
}

func wdmatch(s1, s2 string) string {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	if i == len(s1) {
		return s1
	}
	return ""
}
