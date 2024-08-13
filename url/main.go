package main

import (
	"fmt"
	"net/http"
)

func main() {
	var urls = []string{
		"http://google.com",
		"https://google.com",
		"http://somesite.com",
	}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s адрес не ок\n", url)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("%s-адрес-ok\n", url)
		} else {
			fmt.Printf("%s не ок\n", url)
			return
		}
	}

}
