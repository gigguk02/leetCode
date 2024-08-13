package main

import "fmt"

func loli(strArr []string) []string {
	var a []int
	for i := 0; i < len(strArr); i++ {
		a = append(a, (len(strArr[i]) % 3))
	}
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a); i++ {
			if len(a) > i+1 && a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				strArr[i], strArr[i+1] = strArr[i+1], strArr[i]
			}
		}

	}
	fmt.Println(a)
	fmt.Println(strArr)
	return []string{}

}
func main() {
	fmt.Println(loli([]string{"Colorado", "Utah", "Wisconsin", "Oregon"}))

}
