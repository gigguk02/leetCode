package main

import "fmt"

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
}
func findCenter(edges [][]int) int {
	num := make(map[int]int)
	for _, i := range edges {
		a, b := i[0], i[1]
		num[a]++
		num[b]++
		if num[a] == len(edges) {
			return a
		}
		if num[b] == len(edges) {
			return b
		}
	}

	return 1

}
