package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello %v", name)
}

func Add(chnageQty int, number *int) int {
	*number = *number + chnageQty
	return *number
}

func main() {
	greet("Gopher")

	a := 2
	Add(a, &a)

	fmt.Printf("%v\n", a)
}
