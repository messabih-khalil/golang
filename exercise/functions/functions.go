//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) string{
	return "Hello " + name
}

func message(message string) {
	fmt.Println(message)
}

func addThreeNumbers(num1 , num2 , num3 int) int{
	return num1 + num2 + num3
}

func returnOneNumber() int {
	return 1
}

func returnTwoNumber() (int , int) {
	return 3 , 2
}




func main() {
	// print greeting message 

	greeting := greeting("aldin")

	fmt.Println(greeting)

	// print message

	message("random message")

	// numbers

	num1 := returnOneNumber()

	num2 , num3 := returnTwoNumber()

	// add numbers 

	result := addThreeNumbers(num1 , num2 , num3)

	fmt.Println("Result : " , result)
}
