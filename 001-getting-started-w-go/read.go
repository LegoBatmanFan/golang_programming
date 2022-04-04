// Lena Horsley
// Getting Started with Go
// Peer-graded Assignment: Module 2 Activity: findian.go
//
// 28 June 2021
//
// Write a program which reads information from a file and represents it in a slice of
// structs. Assume that there is a text file which contains a series of names. Each line
// of the text file has a first name and a last name, in that order, separated by a single
// space on the line.

// Your program will define a name struct which has two fields, fname for the first name,
// and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will
// successively read each line of the text file and create a struct which contains the
// first and last names found in the file. Each struct created will be added to a slice,
// and after all lines have been read from the file, your program will have a slice
// containing one struct for each line in the file. After reading all lines from the file,
// your program should iterate through your slice of structs and print the first and last
// names found in each struct.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Constants used in the progam
const (
	MAX_FIELD_SIZE      int    = 20
	MY_PROMPT           string = "Give me a filename: "
	ERROR_PROMPT        string = "ERROR: "
	PRINT_SLICE_MESSAGE string = "Printing the name data..."
)

type PersonalData struct {
	firstName string
	lastName  string
}

func main() {

	// Get the filename from the user
	fileName := getFileNameFromUser()

	// Open the file
	nameDataFromFile, errorMessage := os.Open(fileName)

	// If the file doesn;t exist, throw an error
	if errorMessage != nil {
		fmt.Println(ERROR_PROMPT, errorMessage)
		nameDataFromFile.Close()

		// If the file exits, read the data, create the struct, add to the slice.
		// Once all of the file data has been added to the slice, print the slice.
	} else {
		nameDataSlice := createStructAndAddToSlice(nameDataFromFile)
		nameDataFromFile.Close()
		printMySlice(nameDataSlice)
	}
}

func getFileNameFromUser() string {
	fmt.Print(MY_PROMPT)
	userInputScanner := bufio.NewScanner(os.Stdin)
	userInputScanner.Scan()
	myFileName := userInputScanner.Text()

	return myFileName
}

func createStructAndAddToSlice(myFileName io.Reader) []PersonalData {
	var fullName PersonalData
	mySlice := make([]PersonalData, 0)

	scanner := bufio.NewScanner(myFileName)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineWithFullName := strings.Split(scanner.Text(), " ")
		fullName.firstName = truncateString(lineWithFullName[0])
		fullName.lastName = truncateString(lineWithFullName[1])
		mySlice = append(mySlice, fullName)
	}

	return mySlice
}

// Makes the field size 20 characters
func truncateString(fullString string) string {
	runeString := []rune(fullString)
	return string(runeString[:MAX_FIELD_SIZE])
}

func printMySlice(mySlice []PersonalData) {
	fmt.Println(" ")
	fmt.Println(PRINT_SLICE_MESSAGE)
	for _, myName := range mySlice {
		fmt.Println(myName.firstName + "     " + myName.lastName)
	}
}
