package main

import "fmt"

func printMessage(message string) string{
	return message
}

func sumNum(num1 int , num2 int) int{
	return num1 + num2
}


func main() {
	message := "hello world"

	outputMessage := printMessage(message)

	fmt.Println(outputMessage)

	var (
		num1 = 20
		num2 = 30
	)

	resultSum := sumNum(num1 , num2)

	fmt.Println("sum result is : " , resultSum)
}
