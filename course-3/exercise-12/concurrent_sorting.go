package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately
equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which
sorts ¼ of the array should print the subarray that it will sort. When sorting is complete,
the main goroutine should print the entire sorted list.
*/

func splitAndSortList(intList []int, chunkSize float64, chunkLength int, i int, wg *sync.WaitGroup, c chan []int) {
	listN := intList[i*chunkLength : (i*chunkLength)+chunkLength]
	sort.Ints(listN)
	c <- listN
	fmt.Println("Sorted Chunk: n -", i, " List", listN)
	wg.Done()
}

func stringToInt(strList string, chunkSize float64) []int {
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
	fmt.Println("Enter a sequence of integers: ")
	reader := bufio.NewReader(os.Stdin)
	numList, _ := reader.ReadString('\n')

	chunkSize := 4.0
	intList := stringToInt(numList, chunkSize)
	floatListLen := float64(len(intList))
	chunkLength := int(math.Ceil(float64(floatListLen / chunkSize)))

	c := make(chan []int, int(chunkSize))
	semiSortedList := []int{}

	var wg sync.WaitGroup

	for i := 0; i < int(chunkSize); i++ {
		wg.Add(1)
		go splitAndSortList(intList, chunkSize, chunkLength, i, &wg, c)
	}
	wg.Wait()
	close(c)
	for l := range c {
		semiSortedList = append(semiSortedList, l...)
	}

	sort.Ints(semiSortedList)
	fmt.Println("Sorting Complete: ", semiSortedList)
}
