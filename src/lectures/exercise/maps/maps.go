//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status displaying:
//   - number of servers
//   - number of servers for each status (Online, Offline, Maintenance, Retired)
func printSrvSts(servermap map[string]int) {
	fmt.Println("Number of servers are:", len(servermap))
	serverStat := make(map[int]int)
	for _, srvSts := range servermap {
		serverStat[srvSts]++
	}
	fmt.Println(serverStat[0], "Servers are online")
	fmt.Println(serverStat[1], "Servers are Offline")
	fmt.Println(serverStat[2], "Servers are Maintenance")
	fmt.Println(serverStat[3], "Servers are Retired")

	fmt.Println()

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	//* Set all of the server statuses to `Online` when creating the map
	servermap := make(map[string]int)
	for _, srvName := range servers {
		servermap[srvName] = Online
	}
	printSrvSts(servermap)
	// * After creating the map, perform the following actions:
	//   - call display server info function
	//   - change server status of `darkstar` to `Retired`
	servermap["darkstar"] = Retired
	//   - change server status of `aiur` to `Offline`
	servermap["aiur"] = Offline
	printSrvSts(servermap)
	//   - call display server info function
	//   - change server status of all servers to `Maintenance`
	for _, srvName := range servers {
		servermap[srvName] = Maintenance
	}
	printSrvSts(servermap)
	//   - call display server info function
}
