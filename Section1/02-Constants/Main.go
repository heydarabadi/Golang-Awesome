package main

import "fmt"

const host = "localhost"
const (
	name     = "pouya"
	lastname = "heydarabadi"
)

func main() {
	fmt.Println(host)

	fmt.Println(name + "-" + lastname)
}
