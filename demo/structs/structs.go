package main

import "fmt"

// passenger struct

type Passenger struct{
	name string
	ticketNumber int
	borderd bool
}

// bus struct
type Bus struct {
	FrontSeat Passenger
}

func main() {
	// passengers
	aldin := Passenger{"aldin" , 1 , false}
	fmt.Println(aldin)

	anes := Passenger{"anes" , 1 , false}
	khalil := Passenger{"khalil" , 3 , true}
	tars := Passenger{"tars" , 5 , false}

	fmt.Println(anes.name)
	fmt.Println(khalil.name)
	fmt.Println(tars.name)

	busSeat := Bus{khalil}

	fmt.Println("ticket number for" ,busSeat.FrontSeat.name, "is" , busSeat.FrontSeat.ticketNumber)

}
