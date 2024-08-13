package main

import "fmt"

func main() {
	s := "test"
	b := []byte(s)
	b[0] = 'T'
	fmt.Println(string(b))

}
