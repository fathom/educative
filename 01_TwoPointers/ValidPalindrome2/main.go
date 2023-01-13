package main

import "fmt"

func main() {
	fmt.Println("Valid Palindrome II")
	testCase := "abvvcdcvba"
	if validPalindrome(testCase) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return isPalindromic(&s, left+1, right) || isPalindromic(&s, left, right-1)
		}
		left++
		right--
	}

	return true
}

func isPalindromic(s *string, left, right int) bool {
	for left < right {
		if (*s)[left] != (*s)[right] {
			return false
		}
		left++
		right--
	}
	return true
}
