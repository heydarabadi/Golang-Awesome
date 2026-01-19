package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0

	waitGroup := sync.WaitGroup{}
	//mutex := sync.Mutex{}

	waitGroup.Add(100000)

	for i := 0; i < 100_000; i++ {
		go func() {
			defer waitGroup.Done()

			// 1- Method one For Mutex
			//mutex.Lock()
			//counter++
			//mutex.Unlock()

			// 2- Method two For Atomic Use in Instructions
			atomic.AddInt64(&counter, +1)
		}()
	}

	waitGroup.Wait()
	fmt.Println(counter)
}
