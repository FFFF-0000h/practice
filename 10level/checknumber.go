package main

import "fmt"

func main() {
	fmt.Println(CheckNum("Hello"))           // false
	fmt.Println(CheckNum("Hello1"))          // true
	fmt.Println(CheckNum("123"))             // true
	fmt.Println(CheckNum(""))                // false (empty string)
	fmt.Println(CheckNum("Hello World"))     // false
	fmt.Println(CheckNum("Hello 123 World")) // true
}

func CheckNum(strr string) bool {
	//This is to store the length of the string that would be passed in the 	argument
	numstring := len(strr)

	//This is to loop/iterate over the length of string by index, from 0, we	want it to start from 0 index that is why we set it to be less than numstring which is the length of the string and increment everytime, since I started with 0, the end of the string is the actual length of the string, if it were that I started the loop from 1, the end of the string would be numstring - 1.
	for i := 0; i < numstring; i++ {
		/*
			strr[i] retrieves the first index character value and not the actual value itself and compares it with the character value of 0 and 9 and not the actual values, 0 and 9;
			Position:   0    1    2    3    4    5    6
			Character:  H    2    E    L    L    O    1
			ASCII code: 72   69   69   76   76   79   49

			It sees 72 for H and compares 72 to the character code of 0 and not 72 >= 0. The character code of 0 and 9 are 48 and 57.
		*/
		if strr[i] >= '0' && strr[i] <= '9' {
			return true
		}
	}
	return false
}
