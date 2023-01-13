package main

import "fmt"

func main() {
	fmt.Println("Happy Number")

	if happyNumber(28) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func happyNumber(num int) bool {
	slow := num
	fast := sumDigits(num)

	for fast != 1 && fast != slow {
		slow = sumDigits(slow)
		fast = sumDigits(sumDigits(fast))
	}

	return fast == 1
}

func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

func sumDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}
