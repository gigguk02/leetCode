package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findClosestNumber([]int{2, -1, 1}))        // Output: 1
	fmt.Println(findClosestNumber([]int{-4, -2, 1, 4, 8})) // Output: 1
	fmt.Println(findClosestNumber([]int{-1, 1, -2, 2}))    // Output: 1
	fmt.Println(findClosestNumber([]int{5, -5, -10, 10}))  // Output: 5
}

func findClosestNumber(nums []int) int {
	closest := nums[0]
	for _, num := range nums {
		if math.Abs(float64(num)) < math.Abs(float64(closest)) ||
			(math.Abs(float64(num)) == math.Abs(float64(closest)) && num > closest) {
			closest = num
		}
	}
	return closest
}
