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

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func countUnique(word string) int {

	characters := []string{}
	// create an array to store all values characters
	for _, ch := range word {

		// convert rune to string
		s := fmt.Sprintf("%c", ch)

		// check if character exists
		k, _ := Find(characters, s)
		if k == -1 {
			// add if yes
			fmt.Println(s)

			characters = append(characters, s)
		} else {
			// skip if no
			continue
		}
		fmt.Println(characters)
		fmt.Println("")
	}
	return len(characters)
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

	for {
		// user has to enter a value
		value := getUserInput()

		k, _ := Find(usedChars, value)
		if k == -1 {
			fmt.Println("Value not found (add)")
			fmt.Println("")
			usedChars = append(usedChars, value)
			fmt.Println("")
		} else {
			fmt.Println("Value already used")
			continue
		}

		if strings.Contains(guess, value) {
			correctGuess = append(correctGuess, value) // if guess is correct store it
			fmt.Printf(" - %s - is a character of guess", value)
		} else {
			fmt.Printf(" - %s - is not character of guess", value)
		}

		fmt.Println("")
		fmt.Print("Used charaters")
		fmt.Println(usedChars)

		fmt.Println("")
		fmt.Print("Correct guesses")
		fmt.Println(correctGuess)
	}

}
