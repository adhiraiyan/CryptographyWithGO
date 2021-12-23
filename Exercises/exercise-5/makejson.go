package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively. Your program should use Marshal() to create a JSON
object from the map, and then your program should print the JSON object.
*/

type jsonD struct {
	Name    string
	Address string
}

func main() {
	// Get Name from User
	fmt.Println("Enter Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	// Get Address from User
	fmt.Println("Enter Address: ")
	addr, _ := reader.ReadString('\n')

	// Create a map and add the "name" and "address"
	idMap := map[string]string{
		"name":    name,
		"address": addr,
	}

	// use Marshal() to create a JSON object from the map.
	idMapbyte, errJSON := json.Marshal(idMap)
	if errJSON != nil {
		fmt.Println(errJSON.Error())
	}

	// print the JSON object
	var jsonStruct jsonD
	json.Unmarshal(idMapbyte, &jsonStruct)
	fmt.Println("Decoded Name and Address: ", jsonStruct)
}
