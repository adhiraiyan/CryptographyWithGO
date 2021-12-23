package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)


/*
Write a program which prompts the user to enter integers and stores
the integers in a sorted slice. The program should be written as a loop.
Before entering the loop, the program should create an empty integer
slice of size (length) 3. During each pass through the loop,
the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints
the contents of the slice in sorted order. The slice must grow in size
to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters
the character ‘X’ instead of an integer.
*/

func main(){
	// create an empty int slice of size 3
	var infiniteSlice = make([]int, 0, 3)

	// create an infinite loop
	for {
		// Prompt the user to enter an integer
		fmt.Println("Enter an Integer: ")
		var strVar string
		_, err := fmt.Scan(&strVar)
		if err != nil {
			fmt.Println(err.Error())
		}
		// break the loop if the user input is 'X'
		breakCond := strings.Compare(strVar, "X") == 0
		if breakCond {
			break
		}

		// else convert the string to int add the int to the slice
		intVar, _ := strconv.Atoi(strVar)
		infiniteSlice = append(infiniteSlice, intVar)
		fmt.Println("Integer added to Slice: ", infiniteSlice)

		// sort the slice and print content in sorted order
		sort.Ints(infiniteSlice)
		fmt.Println("Sorted Slice: ", infiniteSlice)

		// slice must grow to accommodate any num of integers the user enters
	}

}
