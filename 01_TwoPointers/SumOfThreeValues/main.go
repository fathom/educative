package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sum of Three Values")
	testCase := []int{3, 7, 1, 2, 8, 4, 5}
	target := 20
	result := findSumOfThree(testCase, target)

	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	fmt.Println(nums)

	left, right, sum := 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		right = len(nums) - 1

		fmt.Println("--", i, left, right, "--")

		for left < right {
			sum = nums[i] + nums[left] + nums[right]

			fmt.Println(nums[i], nums[left], nums[right], "=", sum)

			if sum == target {
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return false
}
