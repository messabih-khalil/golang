//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	// my favorite color
	var myFavColor string = "Black"
	fmt.Println("my favorite color is : " , myFavColor)
	// compound assignment

	age := 20
	dateOfBirth := "2002/09/17"

	fmt.Println("my age  is : " , age)
	fmt.Println("my date of birth is : " , dateOfBirth)
	

	// age in days

	var ageInDays int

	// calculate age in days
	ageInDays = age * 365

	fmt.Println("my age in date is : " , ageInDays)


}
