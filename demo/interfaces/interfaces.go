package main

import "fmt"


type Player interface{
	Playing()
	Pout()
}

type Footballer string
type Basketballer string


func (f Footballer) Playing(){
	fmt.Println("Playing Football")
}
func (f Footballer) Pout(){
	fmt.Println("Playing Pout")
}

func (b Basketballer) Playing(){
	fmt.Println("Playing BasketBall")
}


func playing(p []Player){
	for _ , v := range p{
		v.Playing()
	}
}

func main() {
	footPlayer := []Player{Footballer("aldin") , Footballer("dijo") , Footballer("cr7") , Footballer("messi")}

	playing(footPlayer)
}
