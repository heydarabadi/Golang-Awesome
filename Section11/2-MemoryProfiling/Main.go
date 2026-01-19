package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func allocateMemory() []*int {
	var data []*int
	for i := 0; i < 1000000; i++ {
		n := i
		data = append(data, &n)
	}
	return data
}
func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Could not create memory profile:", err)
		return
	}
	defer f.Close()
	data := allocateMemory()
	runtime.GC() // Run garbage collection before writing profile
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("Could not write memory profile:", err)
		return
	}
	fmt.Println("Memory profiling done with", len(data), "items.")
}
