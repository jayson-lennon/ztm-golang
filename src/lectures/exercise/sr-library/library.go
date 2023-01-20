//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	name         string
	checkout     bool
	checkouttime string
	returntime   string
}
type Member struct {
	name string
	book Book
}

func main() {
	p := fmt.Println
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	book1 := Book{}
	book2 := Book{}
	book3 := Book{}
	john := Member{"John", book1}
	lib := []Book{book1, book2, book3}
	book1.name = "book1"
	book1.checkout = true
	p(lib)
	p(book1)
	p(book2)
	mems := []Member{john}
	p(mems)
}
