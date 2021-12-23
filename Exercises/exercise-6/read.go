package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text
file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name,
and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will
successively read each line of the text file and create a struct which contains the
first and last names found in the file. Each struct created will be added to a slice,
and after all lines have been read from the file, your program will have a slice containing
one struct for each line in the file. After reading all lines from the file, your program
should iterate through your slice of structs and print the first and last names found in each struct.
*/

// define a struct which has two fields
type nameStruct struct {
	fname string
	lname string
}

func main() {
	// prompt the user to enter the filename
	fmt.Println("Enter File Name: ")
	var fileName string
	fmt.Scan(&fileName)
	
	// create slice for struct
	var nameSlice []nameStruct

	// read each line of text file
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		// create a struct with fname and lname where each field will be a string of size 20
		person := nameStruct{
			fname: text[0],
			lname: text[1],
		}

		// add person to a slice
		nameSlice = append(nameSlice, person)

	}
	file.Close()

	// iterate through the slice of structs and print first name and last name
	for _, v := range nameSlice {
		fmt.Println("First Name: ", v.fname, "\tLast Name: ", v.lname)
	}
}
