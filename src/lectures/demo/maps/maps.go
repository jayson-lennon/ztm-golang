package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 2
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk delted, new list:", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup", cereal, "boxes")
	}
	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "on the shopping list")
}
