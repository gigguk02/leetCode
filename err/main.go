package main

import "fmt"

func main() {
	fmt.Println(handle())
}

func handle() error {
	return fmt.Errorf("Error")
}
