package main

import "fmt"

func main() {
	slice := []string{"saeed", "nirvana", "#"}
	for i, element := range slice {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Printf("    %q", ch)
		}
		fmt.Println()

	}

}
