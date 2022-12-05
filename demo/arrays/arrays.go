package main


type Room struct{
	name string
	reserved bool
}


func isReserved(rooms[4] Room){
	for i := 0; i < len(rooms); i++ {
		room := rooms[i].reserved
		if room {
			print("The room is reserved" , "\n")
		}else{
			print("The room is not reserved" , "\n")
		}
	}
}

func main() {
	var rooms[4] Room
	rooms[0] = Room{"room one" , true}
	rooms[1] = Room{"room two" , false}
	rooms[2] = Room{"room three" , false}
	rooms[3] = Room{"room four" , true}

	isReserved(rooms)
}
