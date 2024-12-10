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
    if (!inputBox.checkValidity() || state != State.ACTIVE) {
      console.log("Invalid");
      return;
    }
    console.log(`Value: ${inputBox.value}`);
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
  let remaining: number;
  switch (state) {
    case State.INIT:
      min = view.getUint8(0);
      max = view.getUint8(1);
      text.innerText = `Guess a number between [${min}, ${max}]`;

      inputBox.min = `${min}`;
      inputBox.max = `${max}`;
      state = State.WAIT_NEXT;
      break;
    case State.WAIT:
      // Handle result of our guess
      switch (view.getUint8(0)) {
        case GuessResult.GUESS_CORRECT:
          text.innerHTML += "<br>";
          text.innerText += `You guessed ${lastGuess} in ${guesses}`;
          ws.close();
          return;
        case GuessResult.GUESS_HIGHER:
          text.innerHTML += "<br>";
          text.innerText += "The number is higher";
          break;
        case GuessResult.GUESS_LOWER:
          text.innerHTML += "<br>";
          text.innerText += "The number is lower";
          break;
      }
      state = State.WAIT_NEXT;
      break;
    case State.WAIT_NEXT:
      // Tell user we have some guesses left
      remaining = view.getUint8(0);
      text.innerHTML += "<br>";
      text.innerText += `You have ${remaining} guess(es) left`;
      state = State.ACTIVE;
      break;
    case State.ACTIVE:
      console.error(`received message in unsupported state (${State[state]})`);
      break;
  }
});