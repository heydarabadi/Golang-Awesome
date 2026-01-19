package main

import (
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{New: func() interface{} {
	fmt.Printf("Create New Instance")
	return make([]byte, 1024)
}}

func ProccessData() {
	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf)

	for index, item := range buf {
		buf[index] = byte(item * 2)
	}
}
func ProccessData2() {
	buf := make([]byte, 1024)

	for index, item := range buf {
		buf[index] = byte(item * 2)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		ProccessData()
	}
}
