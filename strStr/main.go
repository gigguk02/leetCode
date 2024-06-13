package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if len(haystack) >= len(needle)+i && haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1

}
func main() {
	fmt.Println(strStr("sadbutsad kdls jkdjk; kdlk", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("abc", "c"))

}
