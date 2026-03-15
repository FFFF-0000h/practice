package main

import "fmt"

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
    // Check if numbers are within valid range (0-99)
    if from < 0 || from > 99 || to < 0 || to > 99 {
        return "Invalid\n"
    }
    
    result := ""
    
    if from <= to {
        // Count UP (including when from == to)
        for i := from; i <= to; i++ {
            // Add formatted number to result
            result += formatNumber(i)
            
            // Add comma and space if not the last number
            if i < to {
                result += ", "
            }
        }
    } else {
        // Count DOWN
        for i := from; i >= to; i-- {
            // Add formatted number to result
            result += formatNumber(i)
            
            // Add comma and space if not the last number
            if i > to {
                result += ", "
            }
        }
    }
    
    // Add newline at the end
    return result + "\n"
}

// Helper function to format a number with leading zero if needed
func formatNumber(num int) string {
    if num < 10 {
        // Manually create "0X" format
        return string(rune('0' + num/10)) + string(rune('0' + num%10))
    }
    // Convert numbers 10-99 to string manually
    tens := num / 10
    ones := num % 10
    return string(rune('0' + tens)) + string(rune('0' + ones))
}
