//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//

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

var (
	noOfRoll   = 30
	noOfDice   = 2
	noOfSlides = 6
)

//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of side

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= noOfRoll; i++ {
		sum := 0
		for j := 1; j <= noOfDice; j++ {
			roll := rand.Intn(noOfSlides) + 1
			sum += roll
			fmt.Println("roll #", i, " dice# ", j, "diced#", roll)
		}
		fmt.Println("sum of the dice in this roll is", sum)
		switch sum := sum; {
		case sum == 2 && noOfDice == 2:
			fmt.Println("Snake eyes")
		case sum == 7:
			fmt.Println("Lucky 7")
		case sum%2 == 0:
			fmt.Println("even")
		case sum%2 != 0:
			fmt.Println("odd")

		}
		// if s%2 == 0 {
		// 	fmt.Println("even")
		// } else {
		// 	fmt.Println("odd")
		// }

	}

}

//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
