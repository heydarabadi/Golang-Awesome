package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func pow(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for item := range input {
			out <- item * item * 100
		}
		close(out)
	}()

	return out
}

func addValue(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for item := range input {
			out <- item + item + 3
		}
		close(out)
	}()
	return out
}
func main() {
	nums := gen(2, 3, 4)
	squares := square(nums)
	powNumbers := pow(squares)
	addValues := addValue(powNumbers)
	for n := range addValues {
		fmt.Println(n)
	}
}
