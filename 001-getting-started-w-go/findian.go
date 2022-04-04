// Lena Horsley
// Getting Started with Go
// Peer-graded Assignment: Module 2 Activity: findian.go
//
// 26 June 2021
//
// Write a program which prompts the user to enter a string. The program searches
// through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program
// should print “Found!” if the entered string starts with the character ‘i’, ends
// with the character ‘n’, and contains the character ‘a’. The program should print
// “Not Found!” otherwise. The program should not be case-sensitive, so it does not
// matter if the characters are upper-case or lower-case.
//
// Examples: The program should print “Found!” for the following example entered
// strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print
// “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var firstLetter string = "i"
	var lastLetter string = "n"
	var testString string = "a"
	var stringLenth int = 0

	// Get the user input and read it in
	fmt.Print("Give me a string: >>>>>   ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	// Show the user the input
	fmt.Print("You entered: ")
	fmt.Println(userInput)

	// Change the string to a rune and get the length
	runeString := []rune(userInput)
	stringLenth = len(runeString)

	// Perform the checks.
	var checkFirstCharacter bool = checkForLetter(string(runeString[0]), firstLetter)
	var checkLastCharacter bool = checkForLetter(string(runeString[stringLenth-1]), lastLetter)
	var checkUserInputString bool = checkForLetter(userInput, testString)

	// The program prints “Found!” if the entered string starts
	// with ‘i’, ends with ‘n’, and contains ‘a’. The program prints “Not Found!”
	if checkFirstCharacter && checkLastCharacter && checkUserInputString {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

// Function for checking the strings
func checkForLetter(myLetter string, myTestLetter string) bool {
	return strings.Contains(strings.ToLower(myLetter), myTestLetter)
}
