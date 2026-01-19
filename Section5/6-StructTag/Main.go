package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name   string `json:"name"`
	Age    int    `json:"age", omitempty`
	Gender string `json:"gender"`
}

func main() {
	pouya := People{Name: "pouya", Age: 18, Gender: "male"}
	parseJson, _ := json.MarshalIndent(pouya, "", "\t")

	fmt.Printf(string(parseJson))
}
