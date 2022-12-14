package main

import "fmt"


type Player interface{
	Playing()
}

type Footballer string
type Basketballer string


func (f Footballer) Playing(){
	fmt.Println("Playing Football")
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
