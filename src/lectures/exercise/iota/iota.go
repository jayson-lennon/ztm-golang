//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

//* Mathematical operations must be defined as constants using iota
const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

//* Write a receiver function that performs the mathematical operation
//  on two operands
func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("unhandled operation")
}

func main() {
	//* The existing function calls in main() represent the API and cannot be changed
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Operation(Subtract)
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := Operation(Multiply)
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Operation(Divide)
	fmt.Println(div.calculate(100, 2)) // = 50
}
