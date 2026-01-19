package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}

	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return firstName, lastName
}

func main() {

	value, err := divide(2, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("value: %v\n", value)
	}

	firstName, lastName := splitName("pouya heydarabadi")
	fmt.Printf("firstName: %v\n", firstName)
	fmt.Printf("lastName: %v\n", lastName)

}
