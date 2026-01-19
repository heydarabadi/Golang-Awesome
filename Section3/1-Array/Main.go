package main

import "fmt"

func main() {
	var arrayInt [2]int
	arrayInt[0] = 1
	arrayInt[1] = 2

	var arrayInt2 [3]int = [3]int{1, 2, 3}
	fmt.Printf("arrayInt2 : %v\n", arrayInt2)

	arrayInt3 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arrayInt3: %v\n", arrayInt3)

	fmt.Printf("Cap : %v Len : %v \n", cap(arrayInt), len(arrayInt))
	fmt.Printf("Cap : %v Len : %v \n", cap(arrayInt2), len(arrayInt2))
	fmt.Printf("Cap : %v Len : %v \n", cap(arrayInt3), len(arrayInt3))

	for index, item := range arrayInt3 {
		fmt.Printf("Index : %d , Item : %d\n", index, item)
	}

	var matrix [2][2]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4

	fmt.Printf("Matrix : %v\n", matrix)

}
