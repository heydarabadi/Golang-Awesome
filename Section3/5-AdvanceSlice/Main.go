package main

import (
	"fmt"
	"slices"
)

func main() {
	sliceTest := []int{123, 3412, 43, 5436, 356, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	existCheck := slices.Contains(sliceTest, 1)
	fmt.Printf("existCheck: %v\n", existCheck)

	sliceTest2 := []int{21, 23, 26}
	sliceTest = append(sliceTest, sliceTest2...)

	fmt.Printf("sliceTest: %v\n", sliceTest)

	slices.Sort(sliceTest)

	fmt.Printf("sliceTest: %v\n", sliceTest)

}
