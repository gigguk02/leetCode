package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("1111", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1111", "1111"))

}
func addBinary(a string, b string) string {
	var result string
	var zero string
	c := false
	if len(a) > len(b) {
		for i := 0; i < len(a)-len(b); i++ {
			zero += "0"
		}
		b = zero + b

	}
	if len(a) < len(b) {
		for i := 0; i < len(b)-len(a); i++ {
			zero += "0"
		}

		a = zero + a
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' && !c {
			result = "0" + result
			c = true
			continue
		}
		if c && a[i] != b[i] {
			result = "0" + result
			c = true
			continue
		}
		if c && a[i] == '0' && b[i] == '0' {
			result = "1" + result
			c = false
			continue

		}
		if a[i] == '1' && b[i] == '1' && c {
			result = "1" + result
			c = true
			continue
		}
		if a[i] != b[i] {
			result = "1" + result
		}
		if a[i] == '0' && b[i] == '0' {
			result = "0" + result
		}
		c = false

	}
	if c {
		result = "1" + result
	}

	return result
}
