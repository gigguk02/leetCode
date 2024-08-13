package main

import "fmt"

type Animal interface {
	Ani() string
}
type Cat struct {
	weight int
}
type Dog struct {
	weight int
}

func (r Cat) Ani() int {
	return r.weight * 2
}
func (d Dog) Ani() int {

	return d.weight
}

func main() {
	cat := Cat{weight: 3}
	fmt.Println(cat.Ani())
	dog := Dog{weight: 4}
	fmt.Println(dog.Ani())
}
