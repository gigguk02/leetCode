package main

import (
	"fmt"
	"strings"
)

var twenty = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Eighteen", "Nineteen"}
var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var thousands = []string{"", "Thousand", "Million", "Billion"}

func numberToWords(num int) string {
	var result string
	var count int
	for num != 0 {
		result = helper(num%1000) + thousands[count] + " " + result
		count++
		num /= 1000
	}
	return strings.TrimSpace(result)

}
func helper(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return twenty[num] + " "
	} else if num < 100 {
		return tens[num/10] + " " + helper(num%10)
	} else {
		return tens[num/100] + " Hundred " + helper(num%100)
	}

}
func main() {
	fmt.Println(numberToWords(1234567))

}
