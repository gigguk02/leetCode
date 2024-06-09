package main

import "fmt"

func main() {

	fmt.Println(removeDuplicates([]int{6, 6, 7, 8, 9, 8}))

}
func removeDuplicates(nums []int) int {
	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			nums[l] = nums[i]
			l++

		}
		prev = nums[i]

	}
	return l

}
