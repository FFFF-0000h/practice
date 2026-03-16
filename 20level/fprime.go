package main

import "fmt"

func main() {
	fmt.Println(FPrime(225225))
	fmt.Println(FPrime(8333325))
	fmt.Println(FPrime(9539))
	fmt.Println(FPrime(804577))
	fmt.Println(FPrime(42))
	fmt.Println(FPrime(0))
	fmt.Println(FPrime(1))
}

func FPrime(n int) string {
	if n <= 1 {
		return ""
	}
	r := ""
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			if r != "" {
				r += "*"
			}
			r += fmtInt(i)
			n /= i
		}
	}
	if n > 1 {
		if r != "" {
			r += "*"
		}
		r += fmtInt(n)
	}
	return r
}

func fmtInt(x int) string {
	if x == 0 {
		return "0"
	}
	b := []byte{}
	for ; x > 0; x /= 10 {
		b = append([]byte{byte('0' + x%10)}, b...)
	}
	return string(b)
}
