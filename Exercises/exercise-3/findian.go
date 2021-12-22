package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
