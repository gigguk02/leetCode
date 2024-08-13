package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
	count := 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[0])-3; j++ {
			if isMagic(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func isMagic(grid [][]int, i, j int) bool {
	seen := make([]bool, 10)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			num := grid[i+x][j+y]
			if num < 1 || num > 9 || seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	return (grid[i][j]+grid[i][j+1]+grid[i][j+2] == 15) &&
		(grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2] == 15) &&
		(grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2] == 15) &&
		(grid[i][j]+grid[i+1][j]+grid[i+2][j] == 15) &&
		(grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1] == 15) &&
		(grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2] == 15) &&
		(grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] == 15) &&
		(grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] == 15)
}
func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
}
