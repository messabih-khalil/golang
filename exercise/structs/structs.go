//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

// rect coordinates struct

type Rectangle struct{
	height float32
	width float32
}

func calculateArea(rect Rectangle){
	area := rect.height * rect.width
	print("Rectangle Area : " , int(area) , "\n")
}

func calculatePerimeter(rect Rectangle){
	perimeter := (rect.height * 2) + (rect.width * 2)
	print("Rectangle Perimeter : " , int(perimeter) , "\n")
}

func main() {
	rect := Rectangle{10 , 5 }
	// calculate Area
	calculateArea(rect)
	// calculate perimeter
	calculatePerimeter(rect)
}
