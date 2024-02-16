package main

import "fmt"

func main() {
	slice := []string{"I", "Hate", "stupid", "hello", "world", "!!!"}

	for i, element := range slice {
		fmt.Println(i, element, ":")

		for _, ch := range element {
			fmt.Printf("  %q\n", ch)
		}
	}
}
