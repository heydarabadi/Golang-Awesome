package main

import "fmt"

type IAnimal interface {
	Eat()
	Sleep()
	Walk()
}

type IHuman interface {
	IAnimal
	Speak()
	Think()
}

type Person struct {
	FirstName string
}

func (person *Person) Speak() {
	fmt.Printf("hi")
}

func main() {
	ss := Person{
		FirstName: "pouya",
	}

	ss.Speak()
}
