//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:

//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	// * Print integers 1 to 50, except:
	//   - Print "Fizz" if the integer is divisible by 3
	//   - Print "Buzz" if the integer is divisible by 5
	//   - Print "FizzBuzz" if the integer is divisible by both 3 and 5
	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)

		}

	}

}
