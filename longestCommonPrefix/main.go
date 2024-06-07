package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && (len(prefix) > len(strs[i]) || prefix != strs[i][:len(prefix)]) {

			prefix = prefix[:len(prefix)-1]

		}
		if prefix == "" {
			return ""

		}
	}
	return prefix

}
func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
