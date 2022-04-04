// Lena Horsley
// Getting Started with Go
// Peer-graded Assignment: Module 2 Activity: trunc.go
//
// 26 June 2021
//
// Write a program which prompts the user to enter a floating point number and prints the
// integer which is a truncated version of the floating point number that was entered.
// Truncation is the process of removing the digits to the right of the decimal place.

package main

import (
	"fmt"
	"math"
	"strconv"
)

// main function
func main() {

	var userInput string

	// Get the user input
	fmt.Print("Give me a floating point number: >>>>>   ")
	fmt.Scanln(&userInput)

	fmt.Print("You entered: ")
	fmt.Println(userInput)

	// Convert the input to a floating point number
	floatingPointNum, errorMessage := strconv.ParseFloat(userInput, 64)

	// If the string CAN NOT be converted to a floating point number (errorMessage != "<nil>"),
	// return an error.
	if errorMessage != nil {
		fmt.Println("ERROR! >>>>> " + userInput + " is invalid input.")
	} else {
		// If the string CAN be converted into a floating point number, the error message is "<nil>".
		// Print the truncated floating point number.
		truncatedFloatingPointNum := math.Trunc(floatingPointNum)
		fmt.Print(userInput + " truncated (the integer): ")
		fmt.Println(truncatedFloatingPointNum)
	}
}
