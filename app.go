package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var i, j int = 0, 3
var countryData = []string{"ghana", "togo", "benin", "kenya", "uganda"}

const lives = 5

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

func getRandomCountry() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(countryData)
	return countryData[rand.Intn(max-min)]
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

func getCharaterPositions(char string, word string) []int {

	pos := []int{}

	for idx, ch := range word {
		if char == fmt.Sprintf("%c", ch) {
			pos = append(pos, idx)
		}
	}

	return pos
}

func guessedResult(slice []string, word string) string {
	var result strings.Builder

	for i := 0; i < len(word); i++ {
		result.WriteString("_ ")
	}

	return result.String()
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
			characters = append(characters, s)
		} else {
			// skip if no
			continue
		}
	}
	return len(characters)
}

func main() {
	usedChars := []string{}
	correctGuess := []string{}
	incorrectGuess := []string{}
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
	guess := getRandomCountry()

	// count the unique items
	uc := countUnique(guess)

	fmt.Println("Guess the country")
	fmt.Println(guessedResult(correctGuess, guess))

	for {
		// user has to enter a value
		value := getUserInput()

		k, _ := Find(usedChars, value)
		if k == -1 {
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
			incorrectGuess = append(incorrectGuess, value) // if guess is correct store it
			fmt.Printf(" - %s - is not character of guess", value)
		}
		fmt.Println("")
		fmt.Println("")

		fmt.Print("Number of chances left : ")
		fmt.Println(lives - len(incorrectGuess))

		fmt.Print("Incorrect guessed letters")
		fmt.Println(incorrectGuess)

		fmt.Print("Correct guesses letters")
		fmt.Println(correctGuess)

		if len(incorrectGuess) == lives {
			break
		}

		if len(correctGuess) == uc {
			break
		}
	}

	if len(incorrectGuess) == lives {
		fmt.Println("")
		fmt.Println("You died!")
		fmt.Printf("The country was %s", guess)
	}

	if len(correctGuess) == uc {
		fmt.Println("")
		fmt.Println("Well done!")
		fmt.Printf("The country was %s", guess)
	}

}
