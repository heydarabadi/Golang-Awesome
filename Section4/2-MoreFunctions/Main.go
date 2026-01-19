package main

import "fmt"

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}

func funcMain() func() int {
	number := 1
	return func() int {
		number++
		return number
	}
}

func main() {
	factorialNumber := factorial(4)
	fmt.Println(factorialNumber)

	numberForTestFuncToFunc := funcMain()
	fmt.Println(numberForTestFuncToFunc())
	fmt.Println(numberForTestFuncToFunc())
	fmt.Println(numberForTestFuncToFunc())
	fmt.Println(numberForTestFuncToFunc())
	fmt.Println(numberForTestFuncToFunc())
	fmt.Println(numberForTestFuncToFunc())
	fmt.Println(numberForTestFuncToFunc())

	testFunction := func(message string) {
		fmt.Printf(message + "\n")
	}

	testFunction("Hello World")
	testFunction("pouya")
	testFunction("ali")
	testFunction("reza")

}
