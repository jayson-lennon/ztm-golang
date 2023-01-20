//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
type Sectag struct {
	name   string
	active bool
}

// * Create functions to activate and deactivate security tags using pointers
func activate(sectag *Sectag) {
	sectag.active = true
}

func deactivate(sectag *Sectag) {
	sectag.active = false
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Sectag) {
	fmt.Println("deactivating......................")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i])
	}
}

func main() {

	// * Perform the following:
	//   - Create at least 4 items, all with active security tags
	item1 := Sectag{"item1", true}
	item2 := Sectag{"item2", true}
	item3 := Sectag{"item3", true}
	item4 := Sectag{"item4", true}
	// - Store them in a slice or array
	items := []Sectag{item1, item2, item3, item4}
	fmt.Println(items)
	// - Deactivate any one security tag in the array/slice
	deactivate(&items[0])
	fmt.Println(items)
	// - Call the checkout() function to deactivate all tags
	checkout(items)
	// - Print out the array/slice after each change
	fmt.Println(items)

}
