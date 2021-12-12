package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// Same as above, different calling convention
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("Initial:", lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("Occupied:", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot)
}
