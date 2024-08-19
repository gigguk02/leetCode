package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	var result []string
	if len(nums) == 0 {
		return result
	}

	start := nums[0]

	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i-1]))
			}
			if i < len(nums) {
				start = nums[i]
			}
		}
	}

	return result
}
