<script setup lang="ts">
import NumberField from './NumberField.vue'
import Keypad from './Keypad.vue'
</script>

<template>
  <div>
    <NumberField />
    <Keypad />
  </div>
</template>

<style lang="css" scoped>
div {
  background-color: red;
  align-items: center;
}
</style>