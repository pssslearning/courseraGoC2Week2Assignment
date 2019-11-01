package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	fmt.Println("\n------------------------------------------------------------")
	fmt.Println("Please introduce the following initial configutarion values:")
	fmt.Println("------------------------------------------------------------")
	acceleration := getAFloatFromStdin("Acceleration ......: ", false)
	initialVelocity := getAFloatFromStdin("Initial Velocity...: ", true)
	initialDisplacement := getAFloatFromStdin("initialDisplacement: ", false)
	fmt.Printf("Formula: displacement s(t) = ½ (%f * t²) + %f * t + %f\n",
		acceleration, initialVelocity, initialDisplacement)
	time := getAFloatFromStdin("Please enter now a value for 't' (time): ", true)
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("The resulting displacement value s(t) is [%f]\n", fn(time))
}

func getAFloatFromStdin(promptText string, mustBePositive bool) float64 {

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(promptText)
		in.Scan()
		aFloat64, err := strconv.ParseFloat(in.Text(), 64)
		if err != nil {
			fmt.Println("Not a Float parseable value, please try again.")
		} else if mustBePositive && aFloat64 < 0 {
			fmt.Println("Must be a Positive value, please try again.")
		} else {
			return aFloat64
		}
	}
}

// GenDisplaceFn takes three float64 arguments:
// acceleration a, initial velocity vo, and initial displacement so.
// GenDisplaceFn() returns a function which computes displacement as
// a function of time, assuming the given values acceleration,
// initial velocity, and initial displacement.
func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	// s = ½ at^2 + vo * t + so
	fn := func(t float64) float64 {
		return a*math.Pow((t), 2)/float64(2) + v0*t + s0
	}
	return fn
}
