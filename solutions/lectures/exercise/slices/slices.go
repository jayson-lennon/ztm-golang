//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	assemblyLine := make([]Part, 3)
	//  - Create an assembly line having any three parts
	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Sheet"

	fmt.Println("3 parts:")
	showLine(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Washer", "Wheel")
	fmt.Println("\nAdded two parts:")
	showLine(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced:")
	showLine(assemblyLine)
}
