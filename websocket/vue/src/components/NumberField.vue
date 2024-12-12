<template>
  <div class="guess-container">
    <span class="label">
      Guess:
    </span>
    <span id="number">
      <slot>Default Number field</slot>
    </span>
  </div>
</template>

<style scoped>
.guess-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.label {
  text-align: left;
}

#number {
  text-align: right;
  min-width: 3ch;
}

span {
  font-size: 100%;
}
</style>
