//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

// * Create a rectangle structure containing a length and width field
type Rectangle struct {
	length int
	width  int
}

// * Using functions, calculate the area and perimeter of a rectangle,
//   - The functions must use the rectangle structure as the function parameter
func area(rect Rectangle) int {
	return rect.length * rect.width
}

func perimeter(rect Rectangle) int {
	return (rect.width * 2) + (rect.length * 2)
}

// - Print the results to the terminal
func printInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {
	rect := Rectangle{length: 3, width: 7}
	//  - Print the results to the terminal
	printInfo(rect)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.length *= 2
	rect.width *= 2
	//  - Print the new results to the terminal
	printInfo(rect)
}

