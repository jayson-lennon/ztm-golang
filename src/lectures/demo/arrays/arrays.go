package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleaness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "cleaned")
		} else {
			fmt.Println(room.name, "dirty")
		}

	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}
	checkCleaness(rooms)
	fmt.Println("Perfomming cleaning...")
	rooms[0].cleaned = true
	rooms[3].cleaned = true
	checkCleaness(rooms)
}
