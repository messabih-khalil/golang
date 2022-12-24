package main

import "fmt"

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



func main() {
	result , error := divide(15,2)

	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(result)

}
