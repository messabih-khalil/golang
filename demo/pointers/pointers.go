package main

import "fmt"

type Counter struct{
	hits int
}

func info(counter *Counter){
	fmt.Println("Counter data is :" , counter.hits)
}

func increments(counter *Counter){
	counter.hits++
	info(counter)
}

func replace(old *string , new string , counter *Counter){
	oldVal := *old
	*old = new
	increments(counter)
	fmt.Println("Change",oldVal,"to",*old)
}


func main() {
	// define struct
	counter := Counter{}
	counter.hits = 0
	increments(&counter)
	// 
	old := "aladin"
	new := "tars"
	replace(&old , new , &counter)
}
 