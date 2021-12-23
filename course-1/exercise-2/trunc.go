package main

import (
	"fmt"
)

/*
Write a program which prompts the user to enter a floating point number and 
prints the integer which is a truncated version of the floating point 
number that was entered. Truncation is the process of removing the digits 
to the right of the decimal place.
*/


func main() {
	var floatNum float64

	fmt.Printf("Input Floating Point Number:\n")
	_, err := fmt.Scan(&floatNum)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Original Input: %f\n", floatNum)

	var truncFloat int = int(floatNum)
	fmt.Printf("Truncated Floating Point: %d\n", truncFloat)
}
