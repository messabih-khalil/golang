package main

import "fmt"

// show route

func showroute(route [] string){
	print("-----------------------\n")
	for i := 0; i < len(route); i++ {
		el := route[i]
		fmt.Println(el)
	}
}

func main() {

	// slice
	route := []string{"Ai" , "Deep learning" , "Full stack"}
	showroute(route)

	// append element
	route = append(route, "Blockchain developer")
	showroute(route)

	otherRoute := []string{"web3 developer" , "data science"}

	route = append(route, otherRoute...)
	showroute(route)
}
