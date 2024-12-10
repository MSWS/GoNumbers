# Number Guessing Game - Web Socket Version

A simple number guessing game implemented using web sockets.

## Behavior
### Initial Connection
1. The client connects to the server.
1. The server sends two bytes to the client, the lower and upper bounds of the range (inclusive) of numbers to guess.
1. The server then sends another byte representing the number of guesses the client has.
1. The server waits for the client to send a guess.
1. Depending on the guess, the server sends a response to the client with the following rule:

| Result        | Value | Description                          |
| ------------- | ----- | ------------------------------------ |
| GUESS_INVALID | 0     | The guess is not a valid number.     |
| GUESS_LOWER   | 1     | The number is lower than the guess.  |
| GUESS_HIGHER  | 2     | The number is higher than the guess. |
| GUESS_CORRECT | 3     | The guess is correct.                |

6. If the guess is correct, the server terminates the connection. Otherwise, the server decrements the number of guesses and sends the client the number of guesses left.
1. The server then waits for the client to send another guess, repeating the process.
1. If the client runs out of guesses, the server sends the correct number and then terminates the connection.

# Design
## Server
A golang websocket server that implements the behavior above.

## Client
An HTML page orchestrated with TypeScript that connects to the server via websockets and plays the game.