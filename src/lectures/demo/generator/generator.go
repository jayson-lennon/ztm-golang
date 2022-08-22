package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	// buffer up to 3 random numbers
	out := make(chan int, 3)

	go func() {
		// Infinite loop to continually add
		// new random numbers as the buffer
		// gets emptied. This will block
		// once the capacity is reached (3)
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

func main() {
	// seed the RNG so we get new numbers each time
	rand.Seed(time.Now().UnixNano())

	// We are using a global generator. Each generator spawns a new goroutine,
	// so we should re-use this one throughout our program.
	randInt := generateRandInt(1, 100)

	// Get some random integers
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

}
