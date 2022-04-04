// Lena Horsley
// Getting Started with Go
// Peer-graded Assignment: Module 3 Activity: slice.go
//
// 26 June 2021
//
// Write a program which prompts the user to enter integers and stores the integers
// in a sorted slice. The program should be written as a loop. Before entering the
// loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer
// to be added to the slice. The program adds the integer to the slice, sorts the
// slice, and prints the contents of the slice in sorted order. The slice must grow in
// size to accommodate any number of integers which the user decides to enter. The
// program should only quit (exiting the loop) when the user enters the character ‘X’
// instead of an integer.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// Create a slice of length 3
	mySlice := make([]int, 0, 3)

	fmt.Print("Give me a number (or enter X to quit): >>>>>   ")
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	myFlag := myScanner.Text()

	// While the user does not enter "x" or "X":
	// 		- If the input is an integer, add it to the slice, sort the slice, and print
	//			the slice.
	//		- Do nothing if the input IS NOT an integer
	// 		- Ask the user for another number
	for !(strings.EqualFold(myFlag, "x")) {
		flagInteger, flagErrorMessage := strconv.Atoi(myFlag)
		if flagErrorMessage == nil {
			mySlice = append(mySlice, flagInteger)
			sort.Ints(mySlice)
			fmt.Println("Your sorted slice: ", mySlice)
		}

		fmt.Println(" ")
		fmt.Print("Give me a new number (or enter X to quit): >>>>>   ")
		myNewScanner := bufio.NewScanner(os.Stdin)
		myNewScanner.Scan()
		myFlag = myNewScanner.Text()
	}

	// After the user input stops, print the final sorted slice
	fmt.Println("FINAL: ", mySlice)
}
