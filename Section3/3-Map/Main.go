package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"pouya": 20,
		"ali":   19,
		"reza":  18,
		"hasan": 15,
	}

	fmt.Printf("studentGrades: %v\n", studentGrades)

	value, okCheck := studentGrades["pouya"]

	fmt.Printf("Value Exist: %v  value: %v\n", okCheck, value)

	if _, ok := studentGrades["bob"]; ok {
		fmt.Printf("studentGrades: %v\n", studentGrades)
	} else {
		fmt.Printf("Value Cannot Be Found \n")
	}

	delete(studentGrades, "ali")

	valueForAli, okCheckAli := studentGrades["ali"]

	fmt.Printf("Value Exist: %v  value: %v\n", okCheckAli, valueForAli)

	var mapConfigWithMake map[string]bool
	if mapConfigWithMake == nil {
		fmt.Printf("mapConfigWithMake is nil\n")
	}

	mapConfigWithMake2 := make(map[string]bool)
	if mapConfigWithMake2 == nil {
		fmt.Printf("mapConfigWithMake2 is nil\n")
	}

}
