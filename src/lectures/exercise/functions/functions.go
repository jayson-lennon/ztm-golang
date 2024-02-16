// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
package main

import "fmt"

// * Write a function that returns any two numbers
func twoTwos() (int, int) {
	return 2, 2
}

// * Write a function that returns any number
func five() int {
	return 5
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThree(a, b, c int) int {
	return a + b + c
}

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Greetings", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func forgetItAll() string {
	return "Forget about it"
}

func main() {
	greetPerson("Bob")
	fmt.Println(forgetItAll())

	//* Add three numbers together using any combination of the existing functions.
	a, b := twoTwos()
	answer := addThree(five(), a, b)

	//  * Print the result
	//* Call every function at least once
	fmt.Println(answer)
}
