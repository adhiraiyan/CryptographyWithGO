package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func BubbleSortVanilla(intList []int) {
	for i := 0; i < len(intList)-1; i += 1 {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
}

func BubbleSortOdd(intList []int, wg *sync.WaitGroup, c chan []int) {
	for i := 1; i < len(intList)-2; i += 2 {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
	wg.Done()
}

func BubbleSortEven(intList []int, wg *sync.WaitGroup, c chan []int) {
	for i := 0; i < len(intList)-1; i += 2 {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
	wg.Done()
}

func ConcurrentBubbleSort(intList []int, wg *sync.WaitGroup, c chan []int) {
	for i := 0; i < len(intList)-1; i += 1 {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
	wg.Done()
}

func main() {
	/*
		When I do BubbleSortVanilla, it takes roughly 15s for a list of size 100000
		When I do BubbleSortOdd followed by BubbleSortEven, it takes roughly 7s
		When I just do ConcurrentBubbleSort it only takes roughly 1.4s.
		For smaller lists, the odd and even phase seems to work best.
		Can't really explain why the single ConcurrentBubbleSort is better?
		Cause of the overhead in creating the two threads and its also processing the
		same or well half the length of the list.
	*/
	// defer profile.Start(profile.MemProfile).Stop()
	rand.Seed(time.Now().Unix())
	intList := rand.Perm(100000)
	fmt.Println("Read a sequence of", len(intList), "elements")

	c := make(chan []int, len(intList))
	var wg sync.WaitGroup

	start := time.Now()
	for j := 0; j < len(intList)-1; j++ {
		// BubbleSortVanilla(intList) // takes roughly 15s

		// wg.Add(2)
		// go BubbleSortOdd(intList, &wg, c)  // takes roughly 7s
		// go BubbleSortEven(intList, &wg, c)

		wg.Add(1)
		go ConcurrentBubbleSort(intList, &wg, c) // takes roughly 1.4s
	}
	wg.Wait()
	elapsed := time.Since(start)

	// Print the sorted integers
	fmt.Println("Sorted List: ", len(intList), "in", elapsed)
}
