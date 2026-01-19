package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func heavyComputation() {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i % 10
	}
	fmt.Println("Computation done:", sum)
}
func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Could not start CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()
	heavyComputation()
}
