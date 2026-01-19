package main

import "fmt"

func main() {

	// C-Style Loop
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello %d \n", i)
	}

	println(" ------------- ")

	// While-Style Loop
	index := 0
	for index < 5 {
		fmt.Println(index)
		index++
	}

	println(" ------------- ")

	// Infinite-Loop
	counter := 0
	for {
		if counter == 10 {
			break
		}
		fmt.Println(counter)
		counter++
	}

	println(" ------------- ")

	// Skiping-Loop
	counterForSkiping := 1
	for {
		if counterForSkiping%2 == 0 {
			counterForSkiping++
			continue
		}

		fmt.Println(counterForSkiping)
		counterForSkiping++

		if counterForSkiping > 20 {
			break
		}

	}

	println(" ------------- ")

	items := []string{"Pouya", "Ali", "Mostafa", "Milad", "Hossein", "Zahra", "Mahdi", "Ehsan", "Siavash", "Pedram", "Danyal"}
	for index, item := range items {
		fmt.Println(index, item)
	}

	println(" ------------- ")

}
