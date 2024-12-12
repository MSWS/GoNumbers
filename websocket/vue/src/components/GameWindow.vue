<script setup lang="ts">
import Keypad from './Keypad.vue'

function test(num: number) {
  console.log(`submitted: ${num}`)
}
</script>

<template>
  <div>
    <Keypad :onSubmit="test" :max=100 />
  </div>
</template>

<style lang="css" scoped>
div {
  background-color: grey;
  width: 50vmin;
  height: 50vmin;
}
</style>