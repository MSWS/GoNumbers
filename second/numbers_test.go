package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var expectedDefaultCorrect, defaultMaxRange int

func init() {
	expectedDefaultCorrect = 50
	defaultMaxRange = 100
}

func TestCreateGame(t *testing.T) {
	game, err := CreateGame(expectedDefaultCorrect, defaultMaxRange)

	if err != nil {
		t.Fatalf("Failed to properly create game: %v", err)
	}

	if game == nil {
		t.Fatalf("Failed to properly create game: game is nil")
	}
}

func TestCreateGame_WithOutOfBounds(t *testing.T) {
	game, err := CreateGame(defaultMaxRange*2, defaultMaxRange)

	if err == nil {
		t.Fatal("CreateGame did not return an error despite invalid correct")
	}

	if game != nil {
		t.Fatalf("CreateGame did not return nil despite invalid correct")
	}
}

func TestGetGuessesLeft(t *testing.T) {
	game, _ := createTestGame(t)
	left := game.GetGuessesLeft()
	expected := game.maxGuesses
	if left != expected {
		t.Fatalf("GetGuessesLeft returned %d when we expected %d", left, expected)
	}

	t.Run("AfterOne", func(t *testing.T) {
		game.AcceptInput(1)
		left := game.GetGuessesLeft()
		expected := game.maxGuesses - 1
		if left != expected {
			t.Fatalf("GetGuessesLeft returned %d when we expected %d", left, expected)
		}
	})

	t.Run("AfterN", func(t *testing.T) {
		game, _ := createTestGame(t)
		for i := 0; i < game.GetGuessesLeft(); i++ {
			game.AcceptInput(1)
			left := game.GetGuessesLeft()
			expected := game.maxGuesses - i - 1

			if left != expected {
				t.Fatalf("After %d guess(es), GetGuessesLeft returned %d when we expected %d",
					i+1, left, expected)
			}
		}
	})

}

func TestAcceptInput(t *testing.T) {
	t.Run("CorrectGuess", func(t *testing.T) {
		game, _ := createTestGame(t)
		result, err := game.AcceptInput(expectedDefaultCorrect)
		if err != nil {
			t.Fatalf("An error occurred accepting input: %v", err)
		}
		if !result {
			t.Fatal("AcceptInput returned false when we guessed the correct number")
		}
	})

	t.Run("IncorrectGuess", func(t *testing.T) {
		game, _ := createTestGame(t)
		result, err := game.AcceptInput(1)
		if err != nil {
			t.Fatalf("An error occurred accepting input: %v", err)
		}
		if result {
			t.Fatal("AcceptInput returned true when we guessed an incorrect number")
		}
	})

	for _, guess := range []int{0, -1, -100, defaultMaxRange + 1, defaultMaxRange * 100} {
		t.Run(fmt.Sprintf("InvalidGuess (%d)", guess), func(t *testing.T) {
			game, _ := createTestGame(t)
			result, err := game.AcceptInput(guess)
			if err == nil {
				t.Fatal("AcceptInput accepted an invalid guess")
			}
			if result {
				t.Fatalf("AcceptInput returned true when we guessed an invalid number (%d)", guess)
			}
		})
	}

	// Runs through all guesses of a game
	runthrough := func(game *Game) {
		guesses := game.GetGuessesLeft()
		for i := 0; i < guesses; i++ {
			result, err := game.AcceptInput(1)
			if err != nil {
				t.Fatalf("An error occurred accepting input: %v", err)
			}
			if result {
				t.Fatal("AcceptInput returned true when we guessed an incorrect number")
			}
		}
	}

	t.Run("OverflowInput-Incorrect", func(t *testing.T) {
		game, _ := createTestGame(t)
		runthrough(game)

		result, err := game.AcceptInput(1)

		if err == nil {
			t.Fatalf("AcceptInput accepted overflowed input")
		}

		if result {
			t.Fatalf("AcceptInput returned true on overflowed input")
		}
	})

	t.Run("OverflowInput-Correct", func(t *testing.T) {
		game, _ := createTestGame(t)
		runthrough(game)

		result, err := game.AcceptInput(game.correct)

		if err == nil {
			t.Fatalf("AcceptInput accepted overflowed input")
		}

		if result {
			t.Fatalf("AcceptInput returned true on overflowed input")
		}
	})
}

func TestGameOutput(t *testing.T) {
	testContains := func(guess int, comparer string) {
		lowComparer := strings.ToLower(comparer)
		t.Run(fmt.Sprintf("%s (%d)", comparer, guess), func(t *testing.T) {
			game, out := createTestGame(t)
			game.AcceptInput(guess)

			if !strings.Contains(strings.ToLower(out.String()), lowComparer) {
				t.Fatalf("Game did not tell us number is %s when we guessed %d", lowComparer, guess)
			}
		})
	}

	testContains(100, "Lower")
	testContains(99, "Lower")
	testContains(expectedDefaultCorrect+1, "Lower")

	testContains(1, "Higher")
	testContains(2, "Higher")
	testContains(expectedDefaultCorrect-1, "Higher")
}

func TestGetInput(t *testing.T) {
	expected := expectedDefaultCorrect
	input := strings.NewReader(fmt.Sprint(expected, "\n"))

	got := GetInput(input, 1, defaultMaxRange)

	if got != expected {
		t.Fatalf("GetInput returned %d when we expected %d", got, expected)
	}

	t.Run("TooLow", func(t *testing.T) {
		expected := 9
		input := strings.NewReader(fmt.Sprint("3\n", expected, "\n"))

		got := GetInput(input, 5, 10)

		if got != expected {
			t.Fatalf("GetInput returned %d when we expected %d", got, expected)
		}
	})

	t.Run("TooHigh", func(t *testing.T) {
		expected := 6
		input := strings.NewReader(fmt.Sprint("11\n", expected, "\n"))

		got := GetInput(input, 5, 10)

		if got != expected {
			t.Fatalf("GetInput returned %d when we expected %d", got, expected)
		}
	})

	t.Run("InclusiveLow", func(t *testing.T) {
		expected := 5
		input := strings.NewReader(fmt.Sprint(expected, "\n"))

		got := GetInput(input, 5, 10)

		if got != expected {
			t.Fatalf("GetInput returned %d when we expected %d", got, expected)
		}
	})

	t.Run("InclusiveHigh", func(t *testing.T) {
		expected := 10
		input := strings.NewReader(fmt.Sprint(expected, "\n"))

		got := GetInput(input, 5, 10)

		if got != expected {
			t.Fatalf("GetInput returned %d when we expected %d", got, expected)
		}
	})
}

func TestEndGame(t *testing.T) {
	t.Run("ZeroGuesses", func(t *testing.T) {
		game, out := createTestGame(t)
		game.EndGame(true)

		exp := "You guessed 50 in 0 guesses."

		if exp != out.String() {
			t.Fatalf("Expected \"%s\" to match \"%s\"", out.String(), exp)
		}
	})

	t.Run("OneGuess", func(t *testing.T) {
		game, out := createTestGame(t)
		game.AcceptInput(game.correct)
		game.EndGame(true)

		exp := "You guessed 50 in 1 guess."

		if exp != out.String() {
			t.Fatalf("Expected \"%s\" to match \"%s\"", out.String(), exp)
		}
	})

	t.Run("TwoGuesses", func(t *testing.T) {
		game, out := createTestGame(t)
		game.AcceptInput(game.correct)
		game.AcceptInput(game.correct)
		game.EndGame(true)

		exp := "You guessed 50 in 2 guesses."

		if exp != out.String() {
			t.Fatalf("Expected \"%s\" to match \"%s\"", out.String(), exp)
		}
	})

	t.Run("CouldNotGuess", func(t *testing.T) {
		game, out := createTestGame(t)
		game.EndGame(false)

		exp := "You could not guess 50."

		if exp != out.String() {
			t.Fatalf("Expected \"%s\" to match \"%s\"", out.String(), exp)
		}
	})
}

func createTestGame(t *testing.T) (*Game, *bytes.Buffer) {
	writer := &bytes.Buffer{}

	game, err := CreateIOGame(
		expectedDefaultCorrect,
		defaultMaxRange,
		writer)

	if err != nil {
		t.Fatalf("Failed to create test game: %v", err)
	}

	if game == nil {
		t.Fatal("Failed to create test game: game is nil")
	}

	return game, writer
}
