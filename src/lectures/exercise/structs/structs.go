//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
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

type Rectangle struct {
	length int
	width  int
}

func perimeter(rect Rectangle) int {
	return 2 * (rect.length + rect.width)
}

func area(rect Rectangle) int {
	return (rect.length * rect.width)
}

func main() {
	rect := Rectangle{4, 5}
	fmt.Println("perimeter is", perimeter(rect))
	fmt.Println("area is", area(rect))

}
