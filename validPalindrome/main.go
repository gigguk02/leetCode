package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var filteredRunes []rune
	for _, r := range s {
		if !unicode.IsSpace(r) && !unicode.IsPunct(r) {
			filteredRunes = append(filteredRunes, r)
		}
	}
	left, right := 0, len(filteredRunes)-1
	for left < right {
		if filteredRunes[left] != filteredRunes[right] {
			return false
		}
		left++
		right--
	}
	return true
}
