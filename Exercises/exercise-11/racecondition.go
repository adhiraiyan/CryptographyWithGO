package main

import (
	"fmt"
)

/*
Race condition occurs when multiple threads or goroutines access the same data in an
unpredictable way.
In this example two functions operate on the same memory address pointed by &i
One of them increments the number, and the other decrements the value by one.

Check for race condition by running: `go run -race racecondition.go`
*/

func decrement(p *int) {
	(*p)--
	fmt.Println(*p)
}

func increment(p *int) {
	(*p)++
	fmt.Println(*p)
}

func main() {
	i := 0
	go increment(&i)
	go decrement(&i)
	i++
	fmt.Println(i)

}
