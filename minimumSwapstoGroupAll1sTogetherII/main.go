package main

import (
	"fmt"
)

func minSwaps(nums []int) int {
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		}
	}

	extendedNums := append(nums, nums...)

	zeroCount := 0
	for i := 0; i < count; i++ {
		if extendedNums[i] == 0 {
			zeroCount++
		}
	}

	minSwaps := zeroCount
	for i := 1; i <= len(nums); i++ {
		if extendedNums[i-1] == 0 {
			zeroCount--
		}
		if extendedNums[i+count-1] == 0 {
			zeroCount++
		}
		if zeroCount < minSwaps {
			minSwaps = zeroCount
		}
	}

	return minSwaps
}

func main() {
	fmt.Println(minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
	fmt.Println(minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
	fmt.Println(minSwaps([]int{1, 1, 0, 0, 1}))
	fmt.Println(minSwaps([]int{1, 0, 0, 0, 0, 0, 1}))
}
