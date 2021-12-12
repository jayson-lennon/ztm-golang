//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	//* Print integers 1 to 50, except:
	for i := 1; i <= 50; i++ {

		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("FizzBuzz")
		} else if divisibleBy3 {
			//  - Print "Fizz" if the integer is divisible by 3
			fmt.Println("Fizz")
		} else if divisibleBy5 {
			//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
