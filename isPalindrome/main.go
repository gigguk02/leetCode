package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(526))
	fmt.Println(isPalindrome(10))

}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	var palindrom string
	for i := len(str) - 1; i >= 0; i-- {
		palindrom += string(str[i])

	}
	return (palindrom == str)
}
