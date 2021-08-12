package main

import (
	"fmt"
)

func main() {
	fmt.Printf(" %v\n",largestPalindrome())
}

func largestPalindrome() int {
	palindrome := 0
	for leftNumber := 999; leftNumber > 100; leftNumber-- {
		for rightNumber := 999; rightNumber > 100; rightNumber-- {
			number := leftNumber * rightNumber
			if isPalindrome(fmt.Sprint(number)) {
				if number > palindrome {
					palindrome = number
				}
			}
		}
	}
	return palindrome
}
func isPalindrome(str string) bool {
	return str == reverse(str)
}
func reverse(str string) (result string) {
    for _, v := range str {
        result = string(v) + result
    }
    return
}