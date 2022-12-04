package main

import "fmt"

// import "fmt"

func main() {
	sum := 0
	// loop
	for i := 0; i < 20; i++ {
		sum += i
	}
	print(sum , "\n")
	// infinite loop !! don't try it
	// for{
	// 	print(0)
	// }
	// while
	
	for sum > 0 {
		fmt.Println(sum)
		sum--
	}
}
