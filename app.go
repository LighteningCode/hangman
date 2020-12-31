package main

import (
	"fmt"
	"strings"
)

var i, j int = 0, 3
var countryData = []string{"ghana", "togo", "benin", "kenya"}

// function that gets user input and returns it
func getUserInput() string {
	var userInput string
	fmt.Println("")
	data, err := fmt.Scanf("%s \n", &userInput)
	if err != nil || data != 1 {
		return "Error"
	}
	return userInput
}

func main() {
	usedChars := []string{}
	correctGuess := []string{}
	fmt.Println("Welcome to hangman, type 's' to begin")

	// get data to start
	data := getUserInput()

	// if returned data is not s keep asking
	for data != "s" {
		data = getUserInput()
	}

	// if data is s
	if data == "s" {
		fmt.Println("Hangman started")
	}

	// choose a the guess
	guess := countryData[2]

	fmt.Println("Guess the country")
	for i := 0; i < len(guess); i++ {
		fmt.Print("_ ")
	}

	// user has to enter a value
	value := getUserInput()
	usedChars = append(usedChars, value)
	if strings.Contains(guess, value) {
		correctGuess = append(correctGuess, value) // if guess is correct store it
		fmt.Printf(" - %s - is a character of guess", value)
	} else {
		fmt.Printf(" - %s - is not character of guess", value)
	}

	fmt.Println("")
	fmt.Println("Used charaters")
	fmt.Println(usedChars)

}
