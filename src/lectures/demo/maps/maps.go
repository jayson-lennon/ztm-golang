package main

import "fmt"

func main() {
	shoppinglist := make(map[string]int)
	shoppinglist["eggs"] = 11
	shoppinglist["milk"] = 1
	shoppinglist["bread"] += 1
	fmt.Println(shoppinglist)
	shoppinglist["eggs"] += 1
	fmt.Println("more eggs!", shoppinglist)
	delete(shoppinglist, "milk")
	fmt.Println("milk deleted", shoppinglist)

	fmt.Println("need", shoppinglist["eggs"], "eggs")
	cereal, found := shoppinglist["cereal"]
	fmt.Println("cereal?")
	if !found {
		fmt.Println("no cereal")
	} else {
		fmt.Println(cereal, "box/es of cereal exist")
	}
	totalitem := 0
	for _, amount := range shoppinglist {
		totalitem += amount
	}
	fmt.Println("Thee are totally", totalitem, "items on the list")

}
