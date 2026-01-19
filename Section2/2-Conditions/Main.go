package main

import "fmt"

func main() {

	age := 20
	if age > 18 {
		fmt.Println("Your Age Gt 18")
	} else if age <= 18 {
		fmt.Println("Your Age lt 18")
	}

	userAccess := map[string]bool{
		"pouya":   true,
		"ali":     false,
		"hossein": true,
		"alireza": false,
		"zahra":   true,
	}

	key := "sas"
	hasAccess := userAccess[key]

	if hasAccess {
		fmt.Printf("%s Has Access To This Object", key)
	} else {
		fmt.Printf("%s Access Not Granted", key)
	}
}
