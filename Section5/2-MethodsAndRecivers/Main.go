package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName  string
	LastName   string
	Age        int
	InsertDate time.Time
}

// Method
func (person *Person) ChangeAge(value int) {
	person.Age += value
}

// Function Or Recivers
func GetName(person Person) string {
	return person.FirstName + "-" + person.LastName
}

func main() {

	siavash := Person{
		FirstName:  "Siavash",
		LastName:   "Ebrahimi",
		Age:        25,
		InsertDate: time.Now(),
	}

	fmt.Printf("siavash: %v \n", siavash)

	siavash.ChangeAge(5)
	fmt.Printf("siavash: %v \n", siavash)

	fullname := GetName(siavash)
	fmt.Printf("FullName: %v \n", fullname)

}
