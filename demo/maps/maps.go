package main

import "fmt"

func main() {
	// create empty map

	myMap := make(map[string]int)

	// insert into map
	myMap["size"] = 20 
	myMap["size2"] = 25
	
	fmt.Println(myMap)

	price , found := myMap["size"]
	if !found {
		fmt.Println("price not found" , price)
	}

	// itteration in map

	for key , value := range myMap{
		fmt.Println(key , value)
	}

	// delete 
	delete(myMap , "size")
}
