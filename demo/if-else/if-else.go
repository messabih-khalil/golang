package main

import "fmt"


func main() {

	// quiz
	quiz1 , quiz2 , quiz3 := 20,15,4

	if quiz1 > quiz2 {
		fmt.Println("quiz one higher then two")
	}else if quiz1 < quiz2 {
		fmt.Println("quiz one lower then two")
	}else if quiz3 < quiz1{
		fmt.Println("quiz one higher then three")
	}

}
