//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

//  - Security tags have two states: active (true) and inactive (false)
type SecurityTag bool

//* Create a structure to store items and their security tag state
type Item struct {
	name string
	tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	purse := Item{"Purse", Active}
	watch := Item{"Watch", Active}

	//  - Store them in a slice or array
	items := []Item{shirt, pants, purse, watch}
	fmt.Println(items)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[0].tag)
	fmt.Println(items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println(items)
}
