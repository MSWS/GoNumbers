package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader

const (
	GUESS_INVALID = iota
	GUESS_LOWER   = iota
	GUESS_HIGHER  = iota
	GUESS_CORRECT
)

type GuessResult int

type Game struct {
	min, max, correct       int
	maxGuesses, guessesUsed int
}

func main() {
	http.HandleFunc("/", CreateNewGame)
}

func (game Game) GetGuessesLeft() int {
	return game.maxGuesses - game.guessesUsed
}

func CreateNewGame(response http.ResponseWriter, request *http.Request) {
	game := Game{
		min:        1,
		max:        100,
		maxGuesses: 8,
	}

	game.Handler(response, request)
}

func (game Game) Handler(response http.ResponseWriter, request *http.Request) {
	ctx, err := upgrader.Upgrade(response, request, nil)

	if err != nil {
		return
	}

	defer ctx.Close()

	game.initializeGame(ctx)

	for {
		if !game.doTurn(ctx) {
			break
		}
	}
}

func (game Game) initializeGame(ctx *websocket.Conn) {
	ctx.WriteMessage(websocket.BinaryMessage, []byte{byte(game.min), byte(game.max)})
}

// Steps the turn of the game
// Returns true if the game should continue (i.e. another turn)
// and false if the game has ended
func (game *Game) doTurn(ctx *websocket.Conn) bool {
	err := ctx.WriteMessage(websocket.BinaryMessage, []byte{byte(game.GetGuessesLeft())})

	if err != nil {
		return false
	}

	mt, msg, err := ctx.ReadMessage()

	if err != nil || mt != websocket.BinaryMessage {
		return false
	}

	guess := msg[0]

	if !game.isValidGuess(int(guess)) {
		ctx.WriteMessage(websocket.BinaryMessage, []byte{GUESS_INVALID})
		return true
	}

	if guess == byte(game.correct) {
		ctx.WriteMessage(websocket.BinaryMessage, []byte{GUESS_CORRECT})
		ctx.Close()
		return false
	}

	if game.GetGuessesLeft() == 1 {
		ctx.Close()
		return false
	}

	game.guessesUsed++

	result := GUESS_HIGHER
	if guess > byte(game.correct) {
		result = GUESS_LOWER
	}
	ctx.WriteMessage(websocket.BinaryMessage, []byte{byte(result)})
	return true
}

func (game Game) isValidGuess(guess int) bool {
	return guess >= game.min && guess <= game.max
}
