package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa"))
}
func isSubsequence(s string, t string) bool {
	var count int
	var k int
	for i := 0; i < len(s); i++ {
		for j := k; j < len(t); j++ {
			if s[i] == t[j] {
				count++
				k = j + 1
				break
			}
		}
	}
	return count == len(s)

}
