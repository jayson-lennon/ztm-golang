package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum is", sum)
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	for sum >= 10 {
		sum -= 5
		fmt.Println("Decrement. Sum is", sum)
	}
}
