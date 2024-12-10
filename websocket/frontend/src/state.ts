export enum State {
  /**
   * We have just initialized and are expecting min + max ranges
   */
  INIT,

  /**
   * We have just sent our guess and are waiting for the server
   * to respond with the result
   */
  WAIT,

  /**
   * We have just received the result of our previous guess,
   * and are waiting for the server to tell us how many guesses
   * we have remaining
   */
  WAIT_NEXT,

  /**
   * We have all info necessary and the server is now waiting
   * for us to send our next guess
   */
  ACTIVE,

  /**
   * We have run out of turns and are waiting for the server
   * to tell us what the number was
   */
  LOST,

  /**
   * The game is no longer active
   */
  INACTIVE
}