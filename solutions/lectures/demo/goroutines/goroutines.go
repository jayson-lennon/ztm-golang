package main

import "fmt"
import "time"
import "unicode"

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	fmt.Printf("Before: %c\n", capitalized)
	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("After: %c\n", capitalized)
}
