//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:

//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

// * Using a slice, create an assembly line that contains type Part
type Part string

// * Create a function to print out the contents of the assembly line
func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func main() {
	assemblyLine := make([]Part, 3)

	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Nut"

	fmt.Println("Three Parts:")
	showLine(assemblyLine)

	assemblyLine = append(assemblyLine, "Washer", "Wheel")
	fmt.Println("\nAdded two Parts:")
	showLine(assemblyLine)

	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced:")
	showLine(assemblyLine)
}
