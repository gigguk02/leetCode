package main

import (
	"fmt"
	"sort"
)

// func minimumPushes(word string) int {
// 	freq := make([]int, 26)
// 	for _, c := range word {
// 		freq[c-'a']++
// 	}
// 	sort.Ints(freq)
// 	totalPushes := 0
// 	multiplier := 1
// 	for i := 25; i >= 0; i-- {
// 		if freq[i] == 0 {
// 			break
// 		}
// 		fmt.Println(i)
// 		if (25-i)%8 == 0 && i != 25 {
// 			multiplier++
// 		}
// 		totalPushes += freq[i] * multiplier
// 	}

//		return totalPushes
//	}
func minimumPushes(word string) int {
	count := make(map[rune]int)
	for _, i := range word {
		count[i]++
	}
	mas := make([]int, 0, len(count))
	for i := range count {
		mas = append(mas, count[i])
	}

	sort.Ints(mas)
	queue := 1
	var result int
	var reverseMas []int
	for i := len(mas) - 1; i >= 0; i-- {
		reverseMas = append(reverseMas, mas[i])

	}
	fmt.Println(reverseMas)
	for i := 0; i < len(reverseMas); i++ {
		if i%8 == 0 && i != 0 {
			queue++
		}
		result += reverseMas[i] * queue
	}
	return result
}
func main() {
	fmt.Println(minimumPushes("abcde"))
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
}
