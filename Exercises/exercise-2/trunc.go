package main

import (
	"fmt"
)

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
