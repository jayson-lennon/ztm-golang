//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array

package main

import "fmt"

type product struct {
	name  string
	price int
}

func counter(list [4]product) {
	sum := 0
	count := 0
	for i := 0; i < len(list); i++ {
		//		fmt.Println(i)
		if list[i].name != "" {
			sum += list[i].price
			count++
		}
	}
	fmt.Println("sum of price is", sum)
	fmt.Println("count of items is", count)
}

func main() {
	list := [4]product{
		{"item1", 10},
		{"item2", 1000},
		{"item3", 23},
	}
	//fmt.Println(len(list))
	//* Print to the terminal:
	//  - The last item on the list
	for i := len(list) - 1; i >= 0; i-- {
		//	fmt.Println(i)
		if list[i].name != "" {
			fmt.Println("last item is", list[i].name)
			break
		}

	}

	//  - The total number of items
	counter(list)
	list[3].name = "item4"
	list[3].price = 500000000
	fmt.Println("Items are updateing ..........")
	counter(list)

	//  - The total cost of the items
	//* Add a fourth product to the list and print out the
	//  information again
	//fmt.Println(list)
}
