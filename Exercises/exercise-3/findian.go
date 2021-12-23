package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to enter a string. 
The program searches through the entered string for the 
characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” 
if the entered string starts with the character ‘i’, ends with the 
character ‘n’, and contains the character ‘a’. The program should 
print “Not Found!” otherwise. The program should not be case-sensitive,
so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example
entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings,
“ihhhhhn”, “ina”, “xian”. 
*/


func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a String:\n")
	enteredString, _ := consoleReader.ReadString('\n')

	fmt.Println("Entered String: ", enteredString)

	// lower case string
	lowString := strings.ToLower(enteredString)
	fmt.Println("LowerCase String: ", lowString)

	// trim any space in string, otherwise indexing didn't seem to work
	trimString := strings.TrimSpace(lowString)
	fmt.Println("Trimmed String: ", trimString)

	cond1 := strings.Compare(string(trimString[0]), "i") == 0
	cond2 := strings.Compare(string(trimString[len(trimString)-1]), "n") == 0
	cond3 := strings.Contains(trimString, "a")
	if (cond1) && (cond2) && (cond3) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
