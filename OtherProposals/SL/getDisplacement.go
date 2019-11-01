package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// GenDisplaceFn takes three float64 arguments, acceleration a, initial velocity v0, and initial displacement s0.
// GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {

	fn := func(t float64) float64 {
		s := 1.0/2.0*a*t*t + v0*t + s0
		return s
	}

	return fn
}

// get an float64 input from stdin
func GetFloat64(prompt string) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		token := scanner.Text()
		if val, err := strconv.ParseFloat(token, 64); err != nil {
			fmt.Println("invalid input: ", token, ", causing error: ", err)
			continue
		} else {
			return val
		}
	}

}

// GetDisplacement calculate the  displacement as function of time t, acceleration a,
// initial velocity vo, and initial displacement so.
func GetDisplacement(a, v0, s0, t float64) float64 {
	return s0 + v0*t + 1.0/2.0*a*t*t
}

func main() {
	a := GetFloat64("Please input acceleration:")

	v0 := GetFloat64("Please input initial velocity:")

	s0 := GetFloat64("Please input initial displacement:")

	fn := GenDisplaceFn(a, v0, s0)

	// prompt the user to enter a value for time
	var t float64
	t = GetFloat64("Please input time:")
	fmt.Printf("t=%v, Calculating with generated function: %v\n", t, fn(t))
	fmt.Printf("t=%v, Direct calculation: %v\n", t, GetDisplacement(a, v0, s0, t))

}
