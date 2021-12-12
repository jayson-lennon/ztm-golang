//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status, including:
func printServerStatus(servers map[string]int) {
	//  - Number of servers
	fmt.Println("\nThere are", len(servers), "servers")

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintenance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Store the existing slice of servers in a map
	serverStatus := make(map[string]int)
	//* Default all of the servers to `Online`
	for _, server := range servers {
		serverStatus[server] = Online
	}

	//* Perform the following status changes and display server info:
	//  - display server info
	printServerStatus(serverStatus)

	//  - change `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired
	//  - change `aiur` to `Offline`
	serverStatus["aiur"] = Offline

	//  - display server info
	printServerStatus(serverStatus)

	//  - change all servers to `Maintenance`
	for server, status := range serverStatus {
		if status == Online {
			serverStatus[server] = Maintenance
		}
	}
	//  - display server info
	printServerStatus(serverStatus)
}
