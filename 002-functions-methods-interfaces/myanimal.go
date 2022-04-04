/*
Write a program which allows the user to create a set of animals and to get information about those
animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user
can either create a new animal of one of the three types, or the user can request information about
an animal that he/she has already created. Each animal has a unique name, defined by the user. Note
that the user can define animals of a chosen type, but the types of animals are restricted to either
cow, bird, or snake. The following table contains the three types of animals and their associated
data.

Animal	Food eaten	Locomtion method	Spoken sound
cow		grass		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a
new prompt on a new line. Your program should continue in this loop forever. Every command from the
user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is
“newanimal”. The second string is an arbitrary string which will be the name of the new animal. The
third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should
process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The
second string is the name of the animal. The third string is the name of the information requested
about the animal, either “eat”, “move”, or “speak”. Your program should process each query command
by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the
Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and
return no values. The Eat() method should print the animal’s food, the Move() method should print
the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three
types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak()
so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an
animal, create an object of the appropriate type. Your program should call the appropriate method
when the user issues a query command.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	INVALID_INPUT_MESSAGE      string = ">>>> INVALID INPUT FOR ANIMAL INFO <<<<<"
	INPUT_PROMPT               string = "Enter animal and info: (CRTL-C to exit the program)"
	RUN_DIVIDER                string = "======================================"
	CANNOT_ADD_EXISTING_ANIMAL string = "ANIMAL EXISTS AND CANNOT BE ADDED"
	INVALID_QUERY              string = "INVALID QUERY: ANIMAL DOES NOT EXIST"
	ANIMAL_CREATED             string = "Created it!"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
	food, locomotion, noise string
}

type Cow struct {
	food, locomotion, noise string
}

type Snake struct {
	food, locomotion, noise string
}

func main() {

	animalMap := make(map[string]Animal)

	for {
		myInput := getAnimalInfoFromUser()
		command := myInput[0]

		if command == "newanimal" {
			animalMap = newAnimal(myInput, animalMap)
		} else if command == "query" {
			animalMap = queryAnimal(myInput, animalMap)
		} else {
			fmt.Println(INVALID_INPUT_MESSAGE)
			fmt.Println(" ")
		}

		fmt.Println(" ")
		fmt.Println(animalMap)
		fmt.Println(RUN_DIVIDER)
		fmt.Println(" ")
	}

}

// Eat(), Move(), Speak() methods for Cow
func (myCow Cow) Eat() {
	fmt.Println(myCow.food)
}

func (myCow Cow) Move() {
	fmt.Println(myCow.locomotion)
}

func (myCow Cow) Speak() {
	fmt.Println(myCow.noise)
}

// Eat(), Move(), Speak() methods for Bird
func (myBird Bird) Eat() {
	fmt.Println(myBird.food)
}

func (myBird Bird) Move() {
	fmt.Println(myBird.locomotion)
}

func (myBird Bird) Speak() {
	fmt.Println(myBird.noise)
}

// Eat(), Move(), Speak() methods for Snake
func (mySnake Snake) Eat() {
	fmt.Println(mySnake.food)
}

func (mySnake Snake) Move() {
	fmt.Println(mySnake.locomotion)
}

func (mySnake Snake) Speak() {
	fmt.Println(mySnake.noise)
}

func getAnimalInfoFromUser() []string {
	var userInputStrings []string
	var userInput2 string = "XXXXXX"

	inputCheck := len(strings.Split(userInput2, " "))
	// Start going through the loop with invalid input. Continue to promt the user until
	// "cow," "bird," or "snake" is entered. Also, the string has to consist of three substrings.
	// Break out of the loop when the coorect input is entered and return the user input strings.
	for inputCheck != 3 && userInput2 != "cow" && userInput2 != "bird" && userInput2 != "snake" && userInput2 != "eat" && userInput2 != "move" && userInput2 != "speak" {
		fmt.Println(INPUT_PROMPT)
		fmt.Print("> ")
		userInput := bufio.NewScanner(os.Stdin)
		userInput.Scan()
		userInputStrings = strings.Split(strings.ToLower(userInput.Text()), " ")
		inputCheck = len(userInputStrings)
		if inputCheck == 3 {
			userInput2 = userInputStrings[2]
		} else {
			fmt.Println(INVALID_INPUT_MESSAGE)
			fmt.Println(" ")
		}
	}

	return userInputStrings
}

func newAnimal(myUserInput []string, myAnimalMap map[string]Animal) map[string]Animal {
	cow := Cow{"grass", "walk", "moo"}
	bird := Bird{"worms", "fly", "peep"}
	snake := Snake{"mice", "slither", "hsss"}

	// check to see if the animal exists
	_, checkFlag := myAnimalMap[myUserInput[1]]

	// If the aminal exists, print a message (because it can't be added). Otherwise, add the animal
	if checkFlag {
		fmt.Println(CANNOT_ADD_EXISTING_ANIMAL)
		fmt.Println(" ")
	} else {
		switch {
		case myUserInput[2] == "cow":
			myAnimalMap[myUserInput[1]] = cow
			fmt.Println(ANIMAL_CREATED)
		case myUserInput[2] == "bird":
			myAnimalMap[myUserInput[1]] = bird
			fmt.Println(ANIMAL_CREATED)
		case myUserInput[2] == "snake":
			myAnimalMap[myUserInput[1]] = snake
			fmt.Println(ANIMAL_CREATED)
		default:
			fmt.Println(INVALID_INPUT_MESSAGE)
			fmt.Println(" ")
		}
	}
	return (myAnimalMap)
}

func queryAnimal(myUserInput []string, myAnimalMap map[string]Animal) map[string]Animal {

	// check to see if the animal exists
	_, checkFlag := myAnimalMap[myUserInput[1]]

	// If the animal does not exist, print a message stating the query request is invalid.
	// Otherwse, perform the query and return the information
	if !checkFlag {
		fmt.Println(INVALID_QUERY)
		fmt.Println(" ")
	} else {
		switch {
		case myUserInput[2] == "speak":
			myAnimalMap[myUserInput[1]].Speak()
		case myUserInput[2] == "move":
			myAnimalMap[myUserInput[1]].Move()
		case myUserInput[2] == "eat":
			myAnimalMap[myUserInput[1]].Eat()
		default:
			fmt.Println(INVALID_INPUT_MESSAGE)
			fmt.Println(" ")
		}
	}
	return myAnimalMap
}
