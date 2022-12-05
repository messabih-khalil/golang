//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServer(servers map[string]int){
	fmt.Println("There are : " , len(servers))
	for key , value := range servers{
		fmt.Println(key , value)
	}
}

// change server status
func changeServerStatus(servers map[string]int){
	for key , _ := range servers{
		server := key
		if server == "darkstar" {
			servers[server] = Retired
		}else if server == "aiur"{
			servers[server] = Offline
		}
	}

	
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	// make server status

	serverStatus := make(map[string]int)

	// Set all of the server statuses to `Online` when creating the map
	for _ , value := range servers{
		serverStatus[value] = Online
	}
	// display server info 
	displayServer(serverStatus)
	// change server status
	changeServerStatus(serverStatus)
	// display server info 
	displayServer(serverStatus)
	//  - change server status of all servers to `Maintenance`
	for key , _ := range serverStatus{
		serverStatus[key] = Maintenance
	}
	// display server info 
	displayServer(serverStatus)
}
