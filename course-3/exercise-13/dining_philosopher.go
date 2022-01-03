package main

import (
	"fmt"
	"sync"
)

/*
Implement the dining philosopher’s problem with the following
constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick
between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we
did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered
first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes
in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it
prints “starting to eat <number>” on a line by itself, where <number> is
the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it
prints “finishing eating <number>” on a line by itself, where <number>
is the number of the philosopher.
*/

// define the Chopstick type
type Chopstick struct{ sync.Mutex }

// define the Philosopher type
type Philosopher struct {
	leftChopStick  *Chopstick
	rightChopStick *Chopstick
	number         int // Constraint 6. Each philosopher is numbered, 1 through 5
	has_eaten      int
}

var wg sync.WaitGroup // a wait group
var messages chan int // channel for messages between goroutines

// Eat method for every philosopher
func (p Philosopher) eat() {
	for {
		// Constraint: each philosopher should eat only 3 times
		if p.has_eaten >= 3 {
			break
		}

		messages <- 1 // add a number to channel to book one space for eating

		// ConstraintIn o: in order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		fmt.Println("Philosopher", p.number, "asking for permission to eat.")

		// Constraint 5. The host allows no more than 2 philosophers to eat concurrently.
		<-messages // proceed if the channel is not full

		p.leftChopStick.Lock()
		p.rightChopStick.Lock()

		// Constraint: when a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
		fmt.Println("Starting to eat", p.number)

		// Constraint: when a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
		p.has_eaten = p.has_eaten + 1 // count how many times he has eaten
		fmt.Println("Finishing eating", p.number)
		if p.has_eaten == 3 {
			fmt.Println("Finishing eating,", p.number, "has eaten", p.has_eaten, "times.")
		} else {
			fmt.Println("Finishing eating", p.number)
		}

		p.rightChopStick.Unlock()
		p.leftChopStick.Unlock()
	}
	wg.Done() // let main goroutine know that this goroutine is complete
}

func main() {

	// Create a channel for messages between goroutines. Constraint: buffer limited to 2 to allow only 2 philosophers to eat at the same time
	messages = make(chan int, 2)

	// Constrain: 5 philosophers share chopsticks, with one chopstick between each adjacent pair of philosophers.
	chopSlice := make([]*Chopstick, 5) // Initialize a slice for 5 chopsticks
	for i := 0; i < 5; i++ {
		chopSlice[i] = new(Chopstick)
	}
	philoSlice := make([]*Philosopher, 5) // Initialize a slice for 5 philosophers
	for i := 0; i < 5; i++ {
		philoSlice[i] = &Philosopher{chopSlice[i], chopSlice[(i+1)%5], i + 1, 0}
	}

	wg.Add(5) // start a wait group and a counter for 5 goroutines

	// run 5 goroutines. Each goroutine will perform the action of eating for each of the 5 philosophers
	for i := 0; i < 5; i++ {
		go philoSlice[i].eat() // concurrently run the eat activity for each philosopher
	}

	wg.Wait() // End of the wait group, decrease the counter

}
