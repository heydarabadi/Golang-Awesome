package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	JoinAt    time.Time
}

func EmployeeFactory(id int, firstName string, lastName string, age int) *Employee {
	return &Employee{Id: id, FirstName: firstName, LastName: lastName, Age: age, JoinAt: time.Now()}
}

func (obj *Employee) Factory(id int, firstName string, lastName string, age int) *Employee {
	return EmployeeFactory(id, firstName, lastName, age)
}

func main() {

	pouya := Employee{
		Id:        1,
		FirstName: "pouya",
		LastName:  "heydarabadi",
		Age:       24,
		JoinAt:    time.Now(),
	}

	hossein := new(Employee)
	hossein.Id = 1
	hossein.FirstName = "hossein"
	hossein.LastName = "rezaei"
	hossein.Age = 25
	hossein.JoinAt = time.Now()

	kimia := &Employee{Id: 1, FirstName: "kimia", LastName: "heydarabadi"}

	parisa := EmployeeFactory(2, "parisa", "rezaei", 28)

	Employee.Factory()

	fmt.Printf("%v \n ", parisa)

	fmt.Printf(" %v \n", kimia)

	fmt.Printf("hossein: %v\n", hossein)

	fmt.Printf("pouya: %v \n", pouya)

}
