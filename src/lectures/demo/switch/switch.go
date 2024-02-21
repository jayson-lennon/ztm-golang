package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap Item")
	case p < 10:
		fmt.Println("Moderately Priced Item")
	default:
		fmt.Println("Expensive Item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy Seating")
	case Business:
		fmt.Println("Business Seating")
	case FirstClass:
		fmt.Println("First Class Seating")
	default:
		fmt.Println("Standing Room Only")
	}
}
