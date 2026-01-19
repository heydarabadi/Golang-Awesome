package main

import (
	"fmt"
	"time"
)

func main() {
	numChannel := make(chan int)

	go SendDataToChannel(numChannel)

	getDataFromChannel1 := <-numChannel
	fmt.Printf("get %v \n", getDataFromChannel1)

	getDataFromChannel2 := <-numChannel
	fmt.Printf("get %v \n", getDataFromChannel2)

	//getDataFromChannel3 := <-numChannel
	//fmt.Printf("%v \n", getDataFromChannel3)

	time.Sleep(5 * time.Second)
}

func SendDataToChannel(numChannel chan int) {

	fmt.Println("before 1")
	numChannel <- 1
	fmt.Println("after 1")

	fmt.Println("before 2")
	numChannel <- 2
	fmt.Println("after 2")

	fmt.Println("before 3")
	numChannel <- 3
	fmt.Println("after 3")
}
