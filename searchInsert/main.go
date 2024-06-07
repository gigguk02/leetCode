package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))

}
func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return i

		}
		if nums[len(nums)-1] < target {
			return i + 1
		}
	}
	return 0

}
