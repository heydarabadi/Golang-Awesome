package main

import "fmt"

func main() {
	names := []string{"pouya", "ali", "hasan"}
	fmt.Printf("names: %v\n", names)

	items := make([]int, 4, 5)
	fmt.Printf("items: %v\n", items)
	fmt.Printf("Cap : %v Len: %v \n", cap(items), len(items))
	for i := 0; i < 20; i++ {
		number := i + 2
		items = append(items, number)
		fmt.Printf("Cap : %v Len: %v \n", cap(items), len(items))

	}

	fmt.Printf("array 3 to 7 : %v \n", items[3:7])
	fmt.Printf("all : %v ", items)

}
