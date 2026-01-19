package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name         string `json:"name"`
	LastName     string `json:"lastName"`
	NationalCode string `json:"nationalCode"`
}

func main() {
	person1 := Person{
		Name:         "Pouya",
		LastName:     "Heydarabadi",
		NationalCode: "231234123123",
	}

	personMarshal, err := json.Marshal(person1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Person marshalling is %s\n", string(personMarshal))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in %v", r)
		}
	}()

	personJson := []byte(`{"name":"Pouya 2","lastName":"Heydarabadi 2 ","nationalCode":"231234123123 2"}`)

	convertToJson := json.Unmarshal(personJson, &person1)
	fmt.Printf("Person marshalling is %s\n", convertToJson)

}
