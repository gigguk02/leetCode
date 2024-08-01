package main

import (
	"fmt"
	"strconv"
)

func countSeniors(details []string) int {
	var count int
	for i := range details {
		a, _ := strconv.Atoi(details[i][11:13])
		if a > 60 {
			count++
		}
	}

	return count

}
func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))

}
