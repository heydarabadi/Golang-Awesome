package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Monday"
	fmt.Println(day)

	switch day {
	case "Sunday", "Monday":
		{
			fmt.Println("ok")
		}
	case "Friday", "Tuesday":
		{
			fmt.Println("Not Ok")
		}
	default:
		{
			fmt.Println("Unknown")
		}

	}

	switch hour := time.Now().Hour(); {
	case hour < 10:
		{
			fmt.Println("good morning")
		}
	case hour > 10:
		{
			fmt.Println("good afternoon")
		}
	case hour > 20:
		{
			fmt.Println("good evening")
		}
	default:
		{
			fmt.Println("good time")
		}
	}
}
