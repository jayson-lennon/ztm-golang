package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A bakers dozen is", bakersDozen)

	moreBakersDozen := add(double(6), 1)
	fmt.Println("More bakers dozen is", moreBakersDozen)
}
