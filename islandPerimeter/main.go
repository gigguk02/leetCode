package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))

}
func islandPerimeter(grid [][]int) int {
	var sum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				sum += 4
				if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
					sum -= 2
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					sum -= 2
				}
			}
		}
	}
	return sum
}
