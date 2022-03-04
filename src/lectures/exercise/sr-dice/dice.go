//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//* The program must handle any number of dice, rolls, and sides
	dice, sides := 2, 12
	rolls := 1

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, ", die #", d, ":", rolled)
		}
		//* Print the sum of the dice roll
		fmt.Println("Total rolled:", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			//  - "Snake eyes": when the total roll is 2, and total dice is 2
			fmt.Println("Snake Eyes!")
		case sum == 7:
			//  - "Lucky 7": when the total roll is 7
			fmt.Println("Lucky Seven!")
		case sum%2 == 0:
			//  - "Even": when the total roll is even
			fmt.Println("Even")
		case sum%2 == 1:
			//  - "Odd": when the total roll is odd
			fmt.Println("Odd")
		}
	}
}
