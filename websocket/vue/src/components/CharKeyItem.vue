<script setup lang="ts">
defineProps<{
  char: string,
  onClick: (char: string) => void
}>()
</script>
<template>
  <button @click="onClick(char)">{{ char }}</button>
</template>

<style lang="css" module>
@import "./KeyCommon.css"
</style>