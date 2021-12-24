package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Each animal has a name and can be either a cow, bird, or snake. With each command,
the user can either create a new animal of one of the three types, or the user can
request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user. Note that the user can define animals
of a chosen type, but the types of animals are restricted to either cow, bird, or snake.

Your program should present the user with a prompt, “>”, to indicate that the user
can type a request. Your program should accept one command at a time from the user,
print out a response, and print out a new prompt on a new line. Your program should
continue in this loop forever. Every command from the user must be either a
“newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings.
The first string is “newanimal”. The second string is an arbitrary string
which will be the name of the new animal. The third string is the type of the new animal,
either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by
creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information
requested about the animal, either “eat”, “move”, or “speak”. Your program should process
each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
which take no arguments and return no values.
The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion, and
the Speak() method should print the animal’s spoken sound.

Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that
the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
*/

type cow struct{ name string }

func (c cow) Eat()   { fmt.Println(c.name, "eats grass") }
func (c cow) Move()  { fmt.Println(c.name, "walk") }
func (c cow) Speak() { fmt.Println(c.name, "speaks moo") }

type bird struct{ name string }

func (b bird) Eat()   { fmt.Println(b.name, "eats worms") }
func (b bird) Move()  { fmt.Println(b.name, "fly") }
func (b bird) Speak() { fmt.Println(b.name, "speaks peep") }

type snake struct{ name string }

func (s snake) Eat()   { fmt.Println(s.name, "eats mice") }
func (s snake) Move()  { fmt.Println(s.name, "moves like slither") }
func (s snake) Speak() { fmt.Println(s.name, "speaks hsss") }

type Animal interface {
	Eat()
	Move()
	Speak()
}

func main() {
	fmt.Println("To Add new animal, command: 'newanimal' Animal Name (valid: cow, bird, snake) Animal Type (valid: cow, bird, snake)")
	fmt.Println("To Query Animal, command: 'query' Animal Name (valid: cow, bird, snake) Request Type (valid: eat, move, speak)")
	fmt.Println("To exit, enter Q")

	reader := bufio.NewReader(os.Stdin)
	objMap := make(map[string]Animal)

	for {
		fmt.Print("> ")
		info, _ := reader.ReadString('\n')

		// break if user enters Q
		breakCond := strings.Compare(strings.TrimSpace(info), "Q") == 0
		if breakCond {
			break
		}

		commands := strings.Split(info, " ")
		if len(commands) != 3 {
			fmt.Println("Incorrect command. Please enter 3 string command")
			break
		}
		switch strings.ToLower(commands[0]) {
		case "newanimal":
			var an Animal
			animalName := strings.TrimSpace(commands[1])
			objectType := strings.TrimSpace(commands[2])
			switch strings.ToLower(objectType) {
			case "cow":
				an = cow{name: animalName}
			case "bird":
				an = bird{name: animalName}
			case "snake":
				an = snake{name: animalName}
			default:
				fmt.Println("error. incorrect object type. allowed types are cow, bird, snake.")
			}
			objMap[animalName] = an
			fmt.Println("Created it!")
		case "query":
			animalName := strings.TrimSpace(commands[1])
			action := strings.TrimSpace(commands[2])
			var animal Animal
			animal, ok := objMap[animalName]
			if !ok {
				fmt.Printf("animal with name %s not present", animalName)
				break
			}
			switch strings.ToLower(action) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("error. incorrect object type. allowed types are cow, bird, snake.")
			}
		default:
			fmt.Println("error - incorrect first command. select b/w newanimal or query.")
		}
	}
}
