//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"


type Cars string
type Trucks string
type Motorcycles string

type Mechanic interface{
	lift()
}

func (car Cars) lift(){
	fmt.Println("Car is the vehicle type,", car, "is a model name")
}
func (truck Trucks) lift(){
	fmt.Println("truck is the vehicle type,", truck, "is a model name")
}
func (moto Motorcycles) lift(){
	fmt.Println("Motorcycles is the vehicle type,", moto, "is a model name")
}

func showInformation(mechanic []Mechanic){
	for _ , v := range mechanic {
		v.lift()
	}
}

func main() {

	listOfVehicle := []Mechanic{Motorcycles("z1000") , Cars("G-class") , Trucks("mercedess")}

	showInformation(listOfVehicle)
}
