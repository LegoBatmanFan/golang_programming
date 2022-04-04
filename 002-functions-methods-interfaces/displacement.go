/*
Lena Horsley
5 July 2021

Let us assume the following formula for displacement s as a function of time t, acceleration a,
initial velocity vo, and initial displacement so.

s = ½ a t^2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity,
and initial displacement. Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time, assuming the given values
acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn()
should take one float64 argument t, representing time, and return one float64 argument which is the
displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity,
and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn()
to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5)
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	ACCELRATION_PROMPT          string = "Enter the acceleration: "
	INITIAL_VELOCITY_PROMPT     string = "Enter the initial velocity: "
	INITIAL_DISPLACEMENT_PROMPT string = "Enter the initial displacement: "
	TIME_PASSED_PROMPT          string = "Enter the time passed: "
	RESULT_PROMPT               string = "Here's your answer: "
	INVALID_INPUT_MESSAGE       string = ">>>>> INAVLID INPUT <<<<<"
)

func main() {

	acceleration := getNumber(ACCELRATION_PROMPT)
	initialVelocity := getNumber(INITIAL_VELOCITY_PROMPT)
	initialDisplacement := getNumber(INITIAL_DISPLACEMENT_PROMPT)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	timePassed := getNumber(TIME_PASSED_PROMPT)
	fmt.Println(RESULT_PROMPT, fn(timePassed))

}

// Calculate the displacment after time t
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	myFunction := func(t float64) float64 {
		return ((a * math.Pow(t, 2)) / 2) + (vo * t) + so
	}

	return myFunction
}

func getNumber(myPrompt string) float64 {
	var userInput string

	// Get the inut from the user
	fmt.Print(myPrompt)
	fmt.Scanln(&userInput)

	// Convert the input to a floating point number
	floatingPointNum, errorMessage := strconv.ParseFloat(userInput, 64)

	// Continue to ask for a number until given valid input
	for errorMessage != nil {
		fmt.Println(INVALID_INPUT_MESSAGE)
		fmt.Print(myPrompt)
		fmt.Scanln(&userInput)
		floatingPointNum, errorMessage = strconv.ParseFloat(userInput, 64)
	}

	return floatingPointNum
}
