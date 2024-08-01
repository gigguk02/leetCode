package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))

}
func isHappy(n int) bool {
	number := make(map[int]bool)
	for !number[n] {
		number[n] = true
		n = sumOfSquares(n)
		if n == 1 {
			return true
		}

	}
	return false

}
func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		q := n % 10
		n = n / 10
		sum += q * q

	}
	return sum

}
