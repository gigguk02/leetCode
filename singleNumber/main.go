package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))

}

func singleNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				break
			}
			return nums[i]

		}
	}
	return 0

}
