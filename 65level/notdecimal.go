package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	j := -1
	n := 0
	if len(dec) == 0 {
		return "\n"
	}
	for i := 0; i < len(dec); i++ {
		if j == -1 && dec[i] == '.' {
			j++
		} else if j == 0 {
			n++
		}
	}
	s, err := strconv.ParseFloat(dec, 64)
	if err == nil {
		return strconv.Itoa(int(s*math.Pow(10, float64(n)))) + "\n"
	}
	return (dec + "\n")
}

/*

func NotDecimal(dec string) string {
 if dec == "" {
 return "\n"
 }
 // Check if valid number and find decimal position
 dotPos := -1
 for i, r := range dec {
 if r == '-' && i == 0 {
 continue
 }
 if r == '.' {
 if dotPos != -1 {
 return dec + "\n" // multiple dots
 }
 dotPos = i
 continue
 }
 if r < '0' || r > '9' {
 return dec + "\n" // not a number
 }
 }
 // No decimal point
 if dotPos == -1 {
 return dec + "\n"
 }
 // Remove decimal point and combine
 result := dec[:dotPos] + dec[dotPos+1:]
 // Remove leading zeros
 if result[0] == '-' {
 for len(result) > 2 && result[1] == '0' {
 result = "-" + result[2:]
 }
 } else {
 for len(result) > 1 && result[0] == '0' {
 result = result[1:]

*/
