package main

import (
	"fmt"
)

func main() {
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
	fmt.Println(findJudge(1, [][]int{}))
}

func findJudge(n int, trust [][]int) int {
	trustCount := make([]int, n+1)
	trustAnyone := make([]bool, n+1)
	for _, t := range trust {
		a, b := t[0], t[1]
		trustCount[b]++
		trustAnyone[a] = true
	}
	for i := 1; i <= n; i++ {
		if !trustAnyone[i] && trustCount[i] == n-1 {
			return i
		}
	}
	return -1
}
