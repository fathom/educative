package main

import "fmt"

func main() {
	fmt.Println("Valid Palindrome")
	if isPalindrome("buzeyzub") {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isPalindrome(inputString string) bool {
	left := 0
	right := len(inputString) - 1

	for left <= right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}

	return true
}
