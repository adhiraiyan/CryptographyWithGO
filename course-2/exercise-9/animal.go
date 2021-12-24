package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats,
2) its method of locomotion, and
3) the sound it makes when it speaks.

Your program should present the user with a prompt, “>”, to indicate that the user can
type a request. Your program accepts one request at a time from the user, prints out the
answer to the request, and prints out a new prompt. Your program should continue in
this loop forever. Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second
string is the name of the information requested about the animal, either “eat”, “move”,
or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion,
and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the
animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.
*/

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (anim Animal) Eat() string {
	return anim.food
}

func (anim Animal) Move() string {
	return anim.locomotion
}

func (anim Animal) Speak() string {
	return anim.noise
}

func main() {

	fmt.Println("Enter an Animal Name (valid: cow, bird, snake), Information Request (valid: eat, move, speak)")
	fmt.Println("To exit, enter Q")

	// this needs to be in an infinite loop
	for {
		// present the user with a prompt ">", accept one request at a time
		fmt.Printf("> ")
		// Request should be a single line with 2 strings, "Animal Name,  Information Request"
		reader := bufio.NewReader(os.Stdin)
		info, _ := reader.ReadString('\n')

		// break if user enters Q
		breakCond := strings.Compare(strings.TrimSpace(info), "Q") == 0
		if breakCond {
			break
		}

		animalList := strings.Split(info, " ")
		animal, action := animalList[0], animalList[1]

		// Initialize the Animal struct with user input
		animalDataMap := map[string]Animal{
			"cow":   {"grass", "walk", "moo"},
			"bird":  {"worms", "fly", "moo"},
			"snake": {"mice", "slither", "hsss"},
		}
		a := animalDataMap[animal]

		// Based on information request, select the appropriate function
		switch strings.TrimSpace(action) {
		case "eat":
			fmt.Println(a.food)
		case "move":
			fmt.Println(a.locomotion)
		case "speak":
			fmt.Println(a.noise)
		default:
			fmt.Println("Invalid Request")
		}
	}
}
