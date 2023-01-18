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
		fmt.Println("Cheap Items")
	case p < 10:
		fmt.Println("Moderately priced Items ")
	default:
		fmt.Println("Moderately priced Items ")
	}
	ticket := Economy
	switch ticket {
	case 0:
		fmt.Println("Economy seating")
	case 1:
		fmt.Println("Business seating")
	case 2:
		fmt.Println("FirstClass seating")
	default:
		fmt.Println("other seating")
	}
}
