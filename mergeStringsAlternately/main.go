package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var result string
	len1, len2 := len(word1), len(word2)

	for i := 0; i < len1 || i < len2; i++ {
		if i < len1 {
			result += string(word1[i])
		}
		if i < len2 {
			result += string(word2[i])
		}
	}
	return result
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
