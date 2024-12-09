package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var correct, guessesLeft, initialGuesses int

func main() {
	initialize()
	gameLoop()
}

// Initialize global vars
func initialize() {
	correct = rand.Intn(100) + 1
	initialGuesses = 6
	guessesLeft = initialGuesses
}

func gameLoop() {
	fmt.Println("Hello! I am thinking of a number [1, 100], what do you think it is?")
	for {
		guessString := "guess"
		if guessesLeft != 1 {
			guessString += "es"
		}
		fmt.Printf("You have %d %s left: ", guessesLeft, guessString)
		if guessesLeft <= 0 || takeTurn() {
			break
		}

		guessesLeft--
	}
}

// Primary game logic for a single turn.
// Returns true if the user got the correct int
func takeTurn() bool {
	guess := getGuess()
	processGuess(guess)

	return guess == correct
}

// Fetches an integer from the user
func getGuess() int {
	var guess int
	lines, err := fmt.Scanf("%d", &guess)

	if lines != 1 {
		fmt.Printf("Hm, I read %d lines, but I can only read one...\n", lines)
		return getGuess()
	}

	if err != nil {
		log.Fatalln(err)
	}

	err = validateGuess(guess)

	if err != nil {
		fmt.Println(err)
		return getGuess()
	}

	return guess
}

func validateGuess(guess int) error {
	if guess <= 0 || guess > 100 {
		return errors.New("guess must be between [1, 100]")
	}

	return nil
}

func processGuess(guess int) {
	if guess == correct {
		fmt.Printf("You guessed my number in %d guesses!\n", initialGuesses-guessesLeft+1)
		os.Exit(0)
		return
	}

	if guessesLeft == 1 {
		return
	}

	var check string

	if guess > correct {
		check = "lower"
	} else {
		check = "higher"
	}

	fmt.Printf("My number is %s\n", check)
}
