package main

import "fmt"

func Summation(numbers ...int) int {

	sum := 0
	for _, item := range numbers {
		sum = sum + item
	}

	return sum
}

func main() {

	sum := Summation(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	fmt.Printf("summation is: %v", sum)

}
