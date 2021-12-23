package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
formula for displacement s as a function of time t, acceleration a, initial velocity v_o,
and initial displacement s_o: s = 1/2 a t^2 + v_o * t + s_o

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial displacement s_o. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t,
representing time, and return one float64 argument which is the displacement travelled after time t.
*/

func stringToFloatList(strList string) []float64 {
	var floatList []float64
	for _, s := range strings.Fields(strList) {
		f, err := strconv.ParseFloat(s, 64)
		if err == nil {
			floatList = append(floatList, f)
		}
	}
	return floatList
}

func GenDisplaceFn(a, v_o, s_o float64) func(float64) float64 {
	displacement := func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v_o * t) + s_o
	}
	return displacement

}

func main() {
	// Get acceleration, initial velocity, and initial displacement from user.
	fmt.Println("Enter Acceleration, Initial Velocity and Initial Displacement space separated: ")
	reader := bufio.NewReader(os.Stdin)
	valList, _ := reader.ReadString('\n')

	// convert the scanned strings to list of integers
	floatList := stringToFloatList(valList)
	a, v_o, s_o := floatList[0], floatList[1], floatList[2]
	fmt.Printf("Read acceleration: %f, Velocity: %f, Displacement: %f", a, v_o, s_o)

	// Prompt the user to enter a value for time
	fmt.Println("\nEnter Time: ")
	var t float64
	fmt.Scan(&t)

	// compute displacement after the entered time
	calcDisplacement := GenDisplaceFn(a, v_o, s_o)
	fmt.Printf("Displacement after time %fs: %f\n", t, calcDisplacement(t))
}
