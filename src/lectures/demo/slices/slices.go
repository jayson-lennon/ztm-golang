package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"home", "work", "gym"}
	printSlice("route1", route)
	route = append(route, "bike shop")
	printSlice("route2", route)
	fmt.Println("home visited")
	fmt.Println("work visited")
	route = route[2:]
	printSlice("new route", route)

}
