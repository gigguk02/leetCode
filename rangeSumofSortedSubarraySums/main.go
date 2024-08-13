package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
}
func rangeSum(nums []int, n int, left int, right int) int {
	mod := 1000000007
	var sum int
	var newMas []int
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum += nums[j]
			newMas = append(newMas, sum)
			
		}
		sum = 0

	}
	sort.Ints(newMas)
	for i := left - 1; i < right; i++ {
		sum += newMas[i] % mod

	}
	return sum

}
