package main

import (
	"fmt"
)

func main() {
	fmt.Println(kthDistinct([]string{"ki", "ci", "vs", "sjoxb"}, 2))

}
func kthDistinct(arr []string, k int) string {
	count := make(map[string]int)
	for _, str := range arr {
		count[str]++
	}
	var unique []string
	for _, str := range arr {
		if count[str] == 1 {
			unique = append(unique, str)
		}
	}
	if k <= len(unique) {
		return unique[k-1]
	}
	return ""
}
