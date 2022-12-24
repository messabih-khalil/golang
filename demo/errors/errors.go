package main

import (
	"errors"
	"fmt"
)

type divNumber struct {
	a , n float64
}

// error message
func (d * divNumber) Error() string {
	return fmt.Sprintf("Cannot divide by zero : %f / %f" , d.a , d.n)
}

// divide function

func divide(a , b float64) (float64 , error) {
	if b == 0 {
		return 0 , &divNumber{a,b}
	}
	return a/b , nil
}

//  -------------

type Stuff struct{
	values []int
}

func (s * Stuff) Get(index int) (int , error) {
	if index > len(s.values) {
		return 0 , errors.New(fmt.Sprintf("The index is not diffiend in the array"))
	}
	return s.values[index] , nil
}

func main() {
	result , error := divide(15,2)

	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(result)

	stuff := Stuff{
	}
	stuff.values = []int{4,5}
	result1 , err := stuff.Get(10)

	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(result1)
}
