package main

import "fmt"

func main() {
	greeting("Saeed")
	fmt.Println(messageRetrun())
	fmt.Println(sumUp(1, 2, 3))
	fmt.Println(RetrunNumber())
	fmt.Println(RetrunTwoNumber())
	a, b := RetrunTwoNumber()
	c := RetrunNumber()
	//fmt.Println("sum of the" ,a , b, c, "is" , sumUp(a, b, c))
	fmt.Println("sum of the", a, b, c, "is", sumUp(a, b, c))

	//fmt.Println(sumUp(RetrunNumber(), RetrunTwoNumber()))
}

// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greeting(name string) {
	fmt.Println("Hello", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func messageRetrun() string {
	return "Some Message"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func sumUp(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func RetrunNumber() int {
	return 2
}

// * Write a function that returns any two numbers
func RetrunTwoNumber() (int, int) {
	return 2, 3
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result

//* Call every function at least once
