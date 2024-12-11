<script setup lang="ts">
defineProps<{
  num: number,
  onClick: (num: number) => void
}>()
</script>
<template>
  <button @click="onClick(num)">{{ num }}</button>
</template>
<style lang="css" module>
@import "./KeyCommon.css"
</style>