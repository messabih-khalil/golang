package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)



func main(){
	msg.Hi()
	display.Display("Hello from main")
	msg.Exciting("I'm exciting")
}