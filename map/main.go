package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	writeToMap := func(key string, value int) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		myMap[key] = value
		fmt.Printf("Записано: %s = %d\n", key, value)
	}
	wg.Add(3)

	go writeToMap("one", 1)
	go writeToMap("two", 2)
	go writeToMap("three", 3)

	wg.Wait()

	mu.Lock()
	for k, v := range myMap {
		fmt.Printf("Ключ: %s, Значение: %d\n", k, v)
	}
	mu.Unlock()
}

// func main() {
// 	myMap := make(map[string]int)
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	writeToMap
// }
