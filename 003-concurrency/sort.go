/*
Lena Horsley

14 July 2021

Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal
size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of
the array should print the subarray that it will sort. When sorting is complete, the main goroutine
should print the entire sorted list.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	FINAL_ARRAY_MESSAGE string = "FINAL ARRAY: "
	INPUT_PROMPT        string = "Enter a list of integers (the integers must be separated by a space) >>>>> "
	SUBARRAY_MESSAGE    string = "here's the sub-array/slice to be sorted: "
	DIVIDER             string = "======================================"
	INPUT_SIZE_MESSAGE  string = "# of integers: "
)

func main() {
	// STEP 1: Get the list of integers from the user
	items := getUserInput()
	var finalNumberArray []int

	// STEP 2. Divide the list into 4 equal parts/arrays/slices
	mySliceOfSlices := createSliceOfSlices(items)
	var myWaitGroup sync.WaitGroup
	fmt.Println(DIVIDER)

	// STEP 3. Create a waitGroup and sort each of the 4 lists with a different goroutine
	myWaitGroup.Add(4)

	go sortSlice(mySliceOfSlices[0], &myWaitGroup)
	go sortSlice(mySliceOfSlices[1], &myWaitGroup)
	go sortSlice(mySliceOfSlices[2], &myWaitGroup)
	go sortSlice(mySliceOfSlices[3], &myWaitGroup)
	myWaitGroup.Wait()

	fmt.Println(DIVIDER)

	// STEP 4. Merge the sorted lists
	for j := 0; j < len(mySliceOfSlices); j++ {
		finalNumberArray = append(finalNumberArray, mySliceOfSlices[j]...)
	}

	sort.Ints(finalNumberArray)

	// STEP 5. Print the final sorted list
	fmt.Println(FINAL_ARRAY_MESSAGE, finalNumberArray)
}

func createSliceOfSlices(items []string) [][]int {
	var count int = 0
	var sliceCounter int = 0
	mySliceOfSlices := make([][]int, 4)

	for count < len(items) {
		if sliceCounter > 3 {
			sliceCounter = 0
		}

		myNumber, _ := strconv.Atoi(items[count])
		mySliceOfSlices[sliceCounter] = append(mySliceOfSlices[sliceCounter], myNumber)
		count++
		sliceCounter++
	}

	//Used for debugging
	fmt.Println(INPUT_SIZE_MESSAGE, count)

	return mySliceOfSlices
}

func getUserInput() []string {

	fmt.Println(INPUT_PROMPT)
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()
	listOfUserNumberStrings := strings.Split(userInput.Text(), " ")

	return listOfUserNumberStrings
}

func sortSlice(myNumberSlice []int, myWaitGroup *sync.WaitGroup) {
	tempSlice := make([]int, 0)

	fmt.Println(SUBARRAY_MESSAGE, myNumberSlice)

	for i := 0; i < len(myNumberSlice); i++ {
		tempSlice = append(tempSlice, myNumberSlice[i])
	}
	sort.Ints(tempSlice)

	if myWaitGroup != nil {
		myWaitGroup.Done()
	}
}
