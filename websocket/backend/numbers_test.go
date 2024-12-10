package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func expectBytes(t *testing.T, conn *websocket.Conn, exp ...byte) {
	mt, msg, err := conn.ReadMessage()

	if err != nil {
		t.Fatalf("An error occurred when attmping to read message: %v", err)
	}

	if mt != websocket.BinaryMessage {
		t.Errorf("Expected message type to be %d, received %d", websocket.BinaryMessage, mt)
	}

	if !bytes.Equal(msg, exp) {
		t.Errorf("Expected byte array to be %s but got %s", hex.EncodeToString(exp), hex.EncodeToString(msg))
	}
}

// Skips initial messages that the websocket will send when we
// start a new game. Our expectted state is start of our turn.
func skipInit(conn *websocket.Conn) {
	conn.ReadMessage() // Read min, max
	conn.ReadMessage() // Read # of turns
}

func TestConnect(t *testing.T) {
	conn, game := setupConnection(t)
	defer conn.Close()

	t.Run("Init", func(t *testing.T) {
		expectBytes(t, conn, byte(game.min), byte(game.max))
	})

	t.Run("TurnBeginning", func(t *testing.T) {
		expectBytes(t, conn, byte(game.maxGuesses))
	})
}

func TestGuess(t *testing.T) {
	check := func(name string, guess byte, exp ...byte) {
		testName := fmt.Sprintf("%s (%d)", name, guess)
		t.Run(testName, func(t *testing.T) {
			conn, _ := setupConnection(t)
			defer conn.Close()

			skipInit(conn)

			conn.WriteMessage(websocket.BinaryMessage, []byte{guess})

			expectBytes(t, conn, exp...)
		})
	}

	check("Invalid", 0, GUESS_INVALID)
	check("Invalid", 255, GUESS_INVALID)
	check("Low", 1, GUESS_HIGHER)
	check("Low", 2, GUESS_HIGHER)
	check("Low", 49, GUESS_HIGHER)
	check("High", 100, GUESS_LOWER)
	check("High", 99, GUESS_LOWER)
	check("High", 51, GUESS_LOWER)
	check("Correct", 50, GUESS_CORRECT)
}

func TestGuessDecrements(t *testing.T) {
	check := func(guess byte, exp ...byte) {
		conn, _ := setupConnection(t)
		defer conn.Close()

		skipInit(conn)

		conn.WriteMessage(websocket.BinaryMessage, []byte{guess})
		conn.ReadMessage() // Skip over result, we only care
		// about the next message that tells us our guesses left
		expectBytes(t, conn, exp...)
	}

	t.Run("High", func(t *testing.T) {
		check(100, 7)
	})

	t.Run("Low", func(t *testing.T) {
		check(1, 7)
	})
}

func setupConnection(t *testing.T) (*websocket.Conn, Game) {
	game := Game{
		min:        1,
		max:        100,
		correct:    50,
		maxGuesses: 8,
	}

	server := *httptest.NewServer(http.HandlerFunc(game.Handler))

	url := fmt.Sprint("ws", strings.TrimPrefix(server.URL, "http"))

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("Failed to create websocket server: %v", err)
	}

	return ws, game
}
