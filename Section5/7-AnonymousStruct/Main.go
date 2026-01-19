package main

import (
	"encoding/json"
	"fmt"
)

type PeopleAnonymouse struct {
	string
	int
}

func main() {

	people := struct {
		FirstName string `json:"people_first_name"`
		LastName  string `json:"people_last_name"`
		Age       int    `json:"people_age"`
	}{
		FirstName: "John",
		LastName:  "Smith",
		Age:       28,
	}

	jsonFormat, _ := json.MarshalIndent(people, "", "\t")
	fmt.Printf(string(jsonFormat) + "\n")

	peopleAnonymousExample := PeopleAnonymouse{
		"pouya",
		26}

	fmt.Printf("People Example %v", peopleAnonymousExample)

}
