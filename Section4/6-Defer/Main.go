package main

import (
	"fmt"
	"os"
)

func test() {
	fmt.Printf("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Printf("4")
}

func main() {

	file, err := os.Create("hi2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	defer func() {
		fmt.Printf("end main")
	}()

	test()

}
