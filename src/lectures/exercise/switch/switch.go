//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func age() int {
	return 15
}

func main() {
	switch a := age(); {
	case a == 0:
		fmt.Println("Newborn")
	case a >= 1 && a <= 3:
		fmt.Println("Toddler")
	case a >= 4 && a <= 12:
		fmt.Println("Child")
	case a >= 13 && a <= 17:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}
}
