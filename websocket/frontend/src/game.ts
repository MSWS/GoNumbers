import { State } from "./state";

const url = "wss://api.msws.dev/numbers";
const text = document.getElementById("outputText")!;
const inputBox = document.getElementById("inputBox") as HTMLInputElement;
const ws = new WebSocket(url);
let min: number, max: number;
let state = State.INIT;

ws.binaryType = "arraybuffer";

inputBox.addEventListener("keydown", (event) => {
  if (event.key === "Enter") {
    ws.send(new Uint8Array([]));
  }
  console.log(min, max);
  console.log(event);
});

ws.addEventListener("message", (ev) => {
  const data = ev.data as Uint8Array;
  switch (state) {
    case State.INIT:
      state = State.WAIT_NEXT;
      min = data[0];
      max = data[1];
      text.textContent = `Guess a number between [${min}, ${max}]`;

      inputBox.min = `${min}`;
      inputBox.max = `${max}`;
      break;
    case State.WAIT:
      // Handle result of our guess
      break;
    case State.WAIT_NEXT:
      // Tell user we have some guesses left
      break;
    case State.ACTIVE:
      console.error("received message");
      break;
    default:
      break;
  }
});