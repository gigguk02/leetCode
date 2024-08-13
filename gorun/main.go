package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	// _,_=<worker(),<-worker()
	fmt.Println(int(time.Since(timeStart).Seconds()))
}
