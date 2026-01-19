package main

import "fmt"

type People struct {
	Name   string
	Age    int
	Gender string
}

type PeopleOption func(*People)

func SetName(name string) PeopleOption {
	return func(People *People) {
		People.Name = name
	}
}

func SetAge(age int) PeopleOption {
	return func(People *People) {
		People.Age = age
	}
}

func SetGender(gender string) PeopleOption {
	return func(People *People) {
		People.Gender = gender
	}
}

func NewPeople(options ...PeopleOption) *People {
	People := new(People)
	for _, option := range options {
		option(People)
	}

	return People
}

func main() {
	People := NewPeople(SetName("James"), SetAge(15), SetGender("Male"))
	fmt.Printf("People is : %v", People)
}
