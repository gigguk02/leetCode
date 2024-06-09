package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(2147395599))

}
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	for i := 1; i < x; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 1

}
