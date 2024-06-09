package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))

}
func lengthOfLastWord(s string) int {
	l := 0
	c := 0
	for i := len(s) - 1; i >= c; i-- {
		if s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
			c = len(s) - 1
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		l++
	}

	return l

}
