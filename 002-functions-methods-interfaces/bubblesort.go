/*
Lena Horsley

4 July 2021

Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up
to 10 integers. The program should print the integers out on one line, in sorted order, from least
to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm
works.

As part of this program, you should write a function called BubbleSort() which takes a slice of
integers as an argument and returns nothing. The BubbleSort() function should modify the slice so
that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of
two adjacent elements in the slice. You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_LIST_LENGTH int = 10
)

func main() {

	listOfUserNumbers := getUserInput()
	numberSlice := convertToSlice(listOfUserNumbers)
	BubbleSort(numberSlice)
	fmt.Println(numberSlice)
}

func Swap(myNumberSlice []int, myIndex int) {
	var temp int = myNumberSlice[myIndex]
	myNumberSlice[myIndex] = myNumberSlice[myIndex+1]
	myNumberSlice[myIndex+1] = temp
}

func BubbleSort(myNumberSlice []int) {
	var myLength int = len(myNumberSlice)

	for i := 0; i < myLength; i++ {
		for j := 0; j < myLength-1; j++ {
			if myNumberSlice[j] > myNumberSlice[j+1] {
				Swap(myNumberSlice, j)
			}
		}
	}
}

func convertToSlice(listOfUserNumberStrings []string) []int {
	numberSlice := make([]int, 0)
	var listLength int = 0

	// If the user entered more than 10 numbers, use the first 10 numbers entered.
	if len(listOfUserNumberStrings) > 10 {
		fmt.Println("# of integers entered: ", len(listOfUserNumberStrings))
		fmt.Println("I can only use the first 10 numbers you entered...")
		listLength = MAX_LIST_LENGTH
	} else {
		listLength = len(listOfUserNumberStrings)
	}

	// Add/append the numbers into a slice
	for i := 0; i < listLength; i++ {
		userNumber, _ := strconv.Atoi(listOfUserNumberStrings[i])
		numberSlice = append(numberSlice, userNumber)
	}
	return numberSlice
}

func getUserInput() []string {
	fmt.Println("Enter 10 integers (the integers must be separated by a space) >>>>> ")
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()
	listOfUserNumberStrings := strings.Split(userInput.Text(), " ")

	return listOfUserNumberStrings
}
