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

import "fmt"
import "time"

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total owned by library
	lended int // total lended
}

//* There must only ever be one copy of the library in memory at any time
type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func returnBook(library *Library, title Title, member *Member) bool {
	// Make sure the book is part of this library
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	// Make sure the member checked out the book
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return true
	}

	// Update library
	book.lended -= 1
	library.books[title] = book

	// Update member info
	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	// Make sure the book is part of the library
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	// Make sure we have enough to lend
	if book.lended == book.total {
		fmt.Println("No more of that book is available")
		return false
	}

	// Update library
	book.lended += 1
	library.books[title] = book

	// Update member info
	member.books[title] = LendAudit{checkOut: time.Now()}

	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}
	//  - Add at least 4 books...
	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The Little Go Book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn Go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}
	//  ... and at least 3 members to the library
	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["Susanna"] = Member{"Susanna", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	//  - Print out initial library information, and after each change
	printLibraryBooks(&library)
	printMemberAudits(&library)

	//  - Check out a book
	member := library.members["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	//  - Check in a book
	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
