package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("Pannic !")
	}
	fmt.Printf("Its ok")
}

func recovable() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in %v", r)
		}
	}()

	mightPanic(true)

}

func main() {
	mightPanic(true)

	recovable()

}
