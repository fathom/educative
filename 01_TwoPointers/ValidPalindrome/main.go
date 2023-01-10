package main

import "fmt"

func main() {
	fmt.Print("Valid Palindrome\n")
	if isPalindrome("buzeyzub") {
		fmt.Print("true\n")
	} else {
		fmt.Print("false\n")
	}
}

func isPalindrome(inputString string) bool {
	// Write your code here
	// Tip: You may use the code template provided
	// in the two_pointers.go file
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
