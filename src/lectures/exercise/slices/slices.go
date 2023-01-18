//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:

package main

import "fmt"

//* Create a function to print out the contents of the assembly line

type Part string

func printAssemblyline(name string, slice []string) {
	fmt.Println("-------", name, "----------")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}

}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	//* Perform the following:
	//  - Create an assembly line having any three parts
	assemblyline := []string{"part1", "part2", "part3"}
	printAssemblyline("line1", assemblyline)
	//  - Add two new parts to the line
	assemblyline = append(assemblyline, "part 4", "part 5")
	printAssemblyline("line 2", assemblyline)
	assemblyline = assemblyline[3:]
	printAssemblyline("line 3", assemblyline)

	//  - Slice the assembly line so it contains only the two new parts
	//  - Print out the contents of the assembly line at each step
	//--Notes:
	//* Your program output should list 3 parts, then 5 parts, then 2 parts

}
