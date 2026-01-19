package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

type PersonBuilder struct {
	Person Person
}

func (personBuilder *PersonBuilder) SetName(name string) *PersonBuilder {
	personBuilder.Person.Name = name
	return personBuilder
}

func (personBuilder *PersonBuilder) SetAge(age int) *PersonBuilder {
	personBuilder.Person.Age = age
	return personBuilder
}

func (personBuilder *PersonBuilder) SetGender(gender string) *PersonBuilder {
	personBuilder.Person.Gender = gender
	return personBuilder
}

func (personBuilder *PersonBuilder) Build() Person {
	return personBuilder.Person
}

func main() {

	personBuilder := new(PersonBuilder)

	pouya := personBuilder.SetName("pouya").SetAge(25).SetGender("Male").Build()
	fmt.Printf("pouya: %v\n", pouya)

	ali := personBuilder.SetName("ali").SetAge(28).SetGender("Male").Build()
	fmt.Printf("ali: %v\n", ali)

}
