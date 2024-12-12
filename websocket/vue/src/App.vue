<script setup lang="ts">
import HeaderItem from './components/HeaderItem.vue'
import GameWindow from './components/GameWindow.vue'
</script>

<template>
  <main>
    <HeaderItem />
    <GameWindow />
  </main>
</template>

<style lang="css" scoped>
main {
  display: flex;
  flex-direction: column;
  align-items: center;
  /* justify-content: center; */
  height: 100vh;
  width: 100vw;


  font-size: xx-large;
}
</style>