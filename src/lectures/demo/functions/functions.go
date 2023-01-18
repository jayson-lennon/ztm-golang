package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lh, rh int) int {
	return lh + rh
}
func greet() {
	fmt.Println("Hello")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("dozein is", dozen)
	bakerdozen := add(dozen, 1)
	fmt.Println("bakerdozein is", bakerdozen)
	anotherbakerdozen := add(double(6), 1)
	fmt.Println("anotherbakerdozen is", anotherbakerdozen)

}
