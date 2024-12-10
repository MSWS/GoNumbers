package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
)

type Game struct {
	correct     int
	guessesUsed int
	maxGuesses  int
	// The min / max possible value for correct, inclusive
	minGuess, maxGuess int
	output             io.Writer
}

func CreateIOGame(correct int, maxGuessRange int, out io.Writer) (*Game, error) {
	if correct > maxGuessRange {
		err := fmt.Errorf("%d cannot be higher than the max guess range of %d", correct, maxGuessRange)
		return nil, err
	}

	game := Game{
		correct:    correct,
		maxGuesses: int(math.Ceil(math.Log2(float64(maxGuessRange)))),
		minGuess:   1,
		maxGuess:   maxGuessRange,
		output:     out,
	}

	return &game, nil
}

func CreateGame(correct int, maxGuessRange int) (*Game, error) {
	return CreateIOGame(correct, maxGuessRange, os.Stdout)
}

func main() {
	game, err := CreateGame(rand.Intn(100)+1, 100)
	fmt.Fprintf(game.output, "I am thinking of a number [%d, %d]...", game.minGuess, game.maxGuess)

	if err != nil {
		log.Fatalf("Unable to create game: %v", err)
	}

	for {
		input := GetInput(os.Stdin, game.minGuess, game.maxGuess)
		result, err := game.AcceptInput(input)
		if err != nil {
			fmt.Fprintf(game.output, "An error occurred when playing: %v", err)
			game.EndGame(false)
		}

		if result {
			game.EndGame(true)
			break
		}

		if game.GetGuessesLeft() <= 0 {
			game.EndGame(false)
			break
		}
	}
}

func GetInput(in io.Reader, min int, max int) int {
	var n int

	for {
		lines, err := fmt.Fscanf(in, "%d", &n)

		if lines != 1 || err != nil {
			continue
		}

		if n < min || n > max {
			continue
		}

		return n
	}
}

func (game *Game) AcceptInput(guess int) (bool, error) {
	if game.GetGuessesLeft() <= 0 {
		return false, errors.New("the game has ended")
	}
	if guess < game.minGuess || guess > game.maxGuess {
		return false, errors.New("invalid guess")
	}
	game.guessesUsed++
	if guess == game.correct {
		return true, nil
	}

	if game.GetGuessesLeft() <= 0 {
		return false, nil
	}

	game.printGuessResult(guess)
	return false, nil
}

func (game *Game) printGuessResult(guess int) {
	comparer := "higher"
	guessStr := "guess"
	guesses := game.GetGuessesLeft()

	if guess > game.correct {
		comparer = "lower"
	}

	if guesses != 1 {
		guessStr += "es"
	}

	fmt.Fprintf(game.output, "My number is %s", comparer)
	fmt.Fprintf(game.output, " You have %d %s left.", guesses, guessStr)
}

func (game *Game) EndGame(won bool) {
	if won {
		str := "guess"
		if game.guessesUsed != 1 {
			str += "es"
		}
		fmt.Fprintf(game.output, "You guessed %d in %d %s.", game.correct, game.guessesUsed, str)

		return
	}

	fmt.Fprintf(game.output, "You could not guess %d.", game.correct)
}

func (game *Game) GetGuessesLeft() int {
	return game.maxGuesses - game.guessesUsed
}
