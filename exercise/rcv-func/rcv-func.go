//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"


type Player struct{
	name string
	health int
	maxHealth int	
	energy int
	maxEnergy int
} 

func (player * Player) modifyEnergy (energy int , maxEnergy int){
	player.energy = energy
	player.maxEnergy = maxEnergy
}

func (player * Player) modifyHealth (health int , maxHealth int){
	player.health = health
	player.maxHealth = maxHealth
}

func (player * Player) showdata(){
	fmt.Println("player name : " , player.name , "player health : " , player.health , "player energy : " , player.energy)
}

func main() {
	// create new player
	player := Player{
		name: "aldin",
		energy: 5,
		maxEnergy: 20,
		health: 7,
		maxHealth: 30,
	}

	player.showdata()

	// modify energy

	fmt.Println("------- modify energy and health -------")

	player.modifyHealth(50,100)
	player.modifyEnergy(10,30)

	player.showdata()
}
