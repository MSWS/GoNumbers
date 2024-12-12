<script setup lang="ts">
defineProps<{
  char: string,
  onClick: (char: string) => void
}>()
</script>
<template>
  <button @click="onClick(char)">{{ char }}</button>
</template>

<style lang="css" scoped>
button {
  border: none;
  border-radius: 20%;
  height: 100%;
  width: 100%;
  font-size: 100%;
}
</style>