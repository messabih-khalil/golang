//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package main

import (
	"fmt"
	"strconv"
	"strings"
)




type Time struct {
	hour , minute , secend int
}

type ErrorMessage struct{
	value string
	message string
}

// handle error message

func (err * ErrorMessage) Error() string{
	return fmt.Sprintf("%v : %v" , err.value , err.message)
}

// convert string into hour , minut , secends

func parseTime(time string) (Time , error) {
	splitString := strings.Split(time, ":")
	if len(splitString) != 3 {
		return Time{} , &ErrorMessage{time , "Invalid String"}
	}else{
		hour , err := strconv.Atoi(splitString[0]) 
		if err != nil {
			return Time{} , &ErrorMessage{time , "Invalid Hour"}
		}
		if hour < 0 || hour > 23 {
			return Time{} , &ErrorMessage{time , "Invalid Hour"}
		}
		minute , err := strconv.Atoi(splitString[1]) 
		if err != nil {
			return Time{} , &ErrorMessage{time , "Invalid minute"}
		}
		if minute < 0 || minute > 59 {
			return Time{} , &ErrorMessage{time , "Invalid minute"}
		}
		secend , err := strconv.Atoi(splitString[2]) 
		if err != nil {
			return Time{} , &ErrorMessage{time , "Invalid secend"}
		}
		if secend < 0 || secend > 59 {
			return Time{} , &ErrorMessage{time , "Invalid secend"}
		}
		return Time{hour , minute , secend} , nil
	}
}

func main(){
	time , err := parseTime("25:00:45")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time)
}