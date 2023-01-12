package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reverse Words in a String")
	testCase := "    a"
	fmt.Println(testCase)
	fmt.Println(reverseWords(testCase))
}

func reverseWords(sentence string) string {
	s := []rune(sentence)
	s = trimString(s)

	s = reverse(s, 0, len(s)-1)

	left := 0
	right := 0
	strLen := len(s) - 1
	for right < strLen && left < strLen {
		if s[left] == ' ' {
			left++
			right = left + 1
			continue
		}

		if s[right] == ' ' && s[right-1] != ' ' {
			s = reverse(s, left, right-1)
			left = right + 1
		}
		right++
	}

	s = reverse(s, left, strLen)

	return string(s)
}

func reverse(s []rune, left int, right int) []rune {
	for left < right {
		x := s[left]
		s[left] = s[right]
		s[right] = x
		left++
		right--
	}
	return s
}

func trimString(s []rune) []rune {
	newS := []rune{}

	i := 0
	j := len(s) - 1

	for {
		if s[i] != ' ' {
			break
		}
		i++
	}

	for {
		if s[j] != ' ' {
			break
		}
		j--
	}

	for i <= j {
		if (s[i] == ' ' && newS[len(newS)-1] != ' ') || s[i] != ' ' {
			newS = append(newS, s[i])
		}
		i++
	}
	return newS
}
