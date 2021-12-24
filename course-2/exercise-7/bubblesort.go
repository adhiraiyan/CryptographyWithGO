package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest.

As part of this program, you should write a function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

func Swap(slice []int, idx int) {
	// Swap the position of two adjacent elements in the slice
	pos0 := slice[0]
	pos1 := slice[1]
	slice[0] = pos1
	slice[1] = pos0
}

func BubbleSort(slice []int) {
	// modify the slice so that the elements are in sorted order
	for j := 0; j < len(slice)-1; j++ {
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice[i:i+2], i)
				// fmt.Println("After Swap idx i: ", i, "slice: ", slice)
			}
			// fmt.Println("idx j: ", j, "slice: ", slice)
		}
	}
}

func stringToInt(strList string) []int {
	var intList []int
	for _, s := range strings.Fields(strList) {
		i, err := strconv.Atoi(s)
		if err == nil {
			intList = append(intList, i)
		}
	}
	return intList
}

func main() {
	// Prompt the user to enter a sequence of 10 integers.
	fmt.Println("Enter a sequence of 10 integers: ")
	reader := bufio.NewReader(os.Stdin)
	numList, _ := reader.ReadString('\n')

	// convert the scanned strings to list of integers
	intList := stringToInt(numList)

	// Sort the list
	BubbleSort(intList)

	// Print the sorted integers
	fmt.Println("Sorted List: ", intList)
}
