package main

import "fmt"

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(f, t int) string {
	if f < 0 || f > 99 || t < 0 || t > 99 {
		return "Invalid\n"
	}

	s, step := "", 1
	if f > t {
		step = -1
	}

	for i := f; ; i += step {
		s += string('0'+i/10) + string('0'+i%10)
		if i == t {
			break
		}
		s += ", "
	}
	return s + "\n"
}

/*

import "strconv"

func FromTo(from, to int) string {
        result := ""

        if from > 99 || from < 0 || to > 99 || to < 0 {
                return "Invalid\n"
        } else if from == to && from < 10 {
                return "0" + strconv.Itoa(from) + "\n"
        }
        if from > to {
                for i := from; i >= to; i-- {
                        if i < 10 {
                                result += "0"
                        }
                        result += strconv.Itoa(i)
                        if i-1 >= to {
                                result += ", "
                        }
                }
                return result + "\n"
        }
        for i := from; i <= to; i++ {
                if i < 10 {
                        result += "0"
                }
                result += strconv.Itoa(i)
                if i+1 <= to {
                        result += ", "
                }
        }
        return result + "\n"
}


*/
