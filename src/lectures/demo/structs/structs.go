package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)
	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)
	bus := Bus{heidi, ella}
	fmt.Println(bus)

}
