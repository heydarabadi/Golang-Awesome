package main

import "fmt"

func main() {
	var number int = 20

	fmt.Printf("number: %v\n", number)

	Calculate(number)
	fmt.Printf("number: %v\n", number)

	Calculate2(&number)
	fmt.Printf("number: %v\n", number)

}

func Calculate(value int) {
	value = value * 10
	fmt.Printf("value: %v\n", value)
}

func Calculate2(value *int) {
	*value = *value * 10
	fmt.Printf("value: %v\n", *value)
}
