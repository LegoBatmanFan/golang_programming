// Lena Horsley
// Getting Started with Go
// Peer-graded Assignment: Module 2 Activity: findian.go
//
// 27 June 2021
//
// Write a program which prompts the user to first enter a name, and then enter an
//address. Your program should create a map and add the name and address to the map
// using the keys “name” and “address”, respectively. Your program should use Marshal()
// to create a JSON object from the map, and then your program should print the JSON
// object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// Get the name and address
	name := getUserInput("Give me a name: >>>>>   ")
	address := getUserInput("Give me an address: >>>>>   ")

	// Create the map and json
	createMapAndJson(name, address)
}

// This function takes the user input and creates the json
func createMapAndJson(myName string, myAddress string) {
	fmt.Println(string(" "))
	// Using the input, create a map
	personalDataMap := map[string]string{"name": myName, "address": myAddress}

	// Now, create the json from the map (if there are no errors)
	myJson, errorMessage := json.Marshal(personalDataMap)
	if errorMessage != nil {
		fmt.Println("ERROR: ", errorMessage)
	} else {
		fmt.Println(string(myJson))
	}
	fmt.Println(string(" "))
}

// This function is used to get the name and the address
func getUserInput(myPrompt string) string {
	fmt.Print(myPrompt)
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	userInput := myScanner.Text()
	return userInput
}
