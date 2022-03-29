package main

import (
	"fmt"
	"strconv"
)

func main() {
	number1 := 17
	number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

	diff := 0
	if number1 < number2 {
		diff = number2 - number1
		resultMessage = "Number2 is greater than Number1 by " + strconv.Itoa(diff)
	} else if number1 > number2 {
		diff = number1 - number2
		resultMessage = "Number1 is greater than Number2 by " + strconv.Itoa(diff)
	} else {
		resultMessage = "Number1 is equal to Number2"
	}
	fmt.Println(resultMessage)

	resultMessage = "Number1(" + (strconv.Itoa(number1)) + ") is "
	if number1%2 == 1 {
		resultMessage = resultMessage + "odd"
	} else {
		resultMessage = resultMessage + "even"
	}
	fmt.Println(resultMessage)

	resultMessage = "Number2(" + (strconv.Itoa(number2)) + ") is "
	if number2%2 == 1 {
		resultMessage = resultMessage + "odd"
	} else {
		resultMessage = resultMessage + "even"
	}
	fmt.Println(resultMessage)
}
