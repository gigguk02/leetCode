package main

import (
	"fmt"
)

func main() {
	rows, cols, rStart, cStart := 3, 4, 1, 2
	result := spiralMatrixIII(rows, cols, rStart, cStart)
	for _, pos := range result {
		fmt.Println(pos)
	}
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	result := [][]int{}
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direction := 0
	steps := 1
	x, y := rStart, cStart

	result = append(result, []int{x, y})

	for len(result) < rows*cols {
		for i := 0; i < steps; i++ {
			x += directions[direction][0]
			y += directions[direction][1]
			if x >= 0 && x < rows && y >= 0 && y < cols {
				result = append(result, []int{x, y})
			}
		}
		direction = (direction + 1) % 4
		if direction%2 == 0 {
			steps++
		}
	}

	return result
}
