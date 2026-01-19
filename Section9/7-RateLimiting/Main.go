package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(500 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("Processing request", req, "at", time.Now())
	}
}
