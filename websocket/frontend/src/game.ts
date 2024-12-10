import { GuessResult } from "./result.js";
import { State } from "./state.js";

const url = "wss://api.msws.dev/numbers";
const text = document.getElementById("outputText")!;
const inputBox = document.getElementById("inputBox") as HTMLInputElement;
const ws = new WebSocket(url);
let min: number, max: number;
let lastGuess: number;
let state = State.INIT;
let guesses = 0;

ws.binaryType = "arraybuffer";

inputBox.addEventListener("keydown", (event) => {
  if (event.key === "Enter") {
    if (!inputBox.checkValidity() || state != State.ACTIVE)
      return;
    if (inputBox.value.length === 0)
      return;

    lastGuess = inputBox.valueAsNumber;

    guesses++;
    state = State.WAIT;
    inputBox.value = "";
    ws.send(new Uint8Array([lastGuess]));
  }
});

ws.addEventListener("message", (ev) => {
  if (!(ev.data instanceof ArrayBuffer)) {
    console.error(`received invalid message from socket: ${ev.data}`);
    return;
  }

  const data = ev.data as ArrayBuffer;
  const view = new DataView(data);
  let remaining, correct: number;
  switch (state) {
    case State.INIT:
      min = view.getUint8(0);
      max = view.getUint8(1);
      text.innerText = `Guess a number between [${min}, ${max}]`;
      text.innerHTML += "<br>";

      inputBox.disabled = false;
      inputBox.min = `${min}`;
      inputBox.max = `${max}`;
      state = State.WAIT_NEXT;
      break;
    case State.WAIT:
      text.innerHTML += "<br>";
      // Handle result of our guess
      switch (view.getUint8(0)) {
        case GuessResult.GUESS_CORRECT:
          text.innerText += `You guessed ${lastGuess} in ${guesses}`;
          ws.close();
          return;
        case GuessResult.GUESS_HIGHER:
          text.innerText += `The number is higher than ${lastGuess}`;
          break;
        case GuessResult.GUESS_LOWER:
          text.innerText += `The number is lower than ${lastGuess}`;
          break;
      }
      text.innerHTML += "<br>";
      state = State.WAIT_NEXT;
      break;
    case State.WAIT_NEXT:
      // Tell user we have some guesses left
      remaining = view.getUint8(0);

      text.innerHTML += "<br>";

      if (remaining === 0) {
        state = State.LOST;
        return;
      }

      text.innerText += `You have ${remaining} guess${remaining == 1 ? "" : "es"} left`;
      state = State.ACTIVE;
      break;
    case State.LOST:
      correct = view.getUint8(0);

      text.innerHTML += "<br>";
      text.innerText += `You failed to guess the number, it was ${correct}`;
      break;
    case State.ACTIVE:
      console.error(`received message in unsupported state (${State[state]})`);
      break;
  }
});