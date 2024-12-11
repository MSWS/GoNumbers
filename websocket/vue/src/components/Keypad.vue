<script setup lang="ts">
import NumKeyItem from './NumKeyItem.vue'
import CharKeyItem from './CharKeyItem.vue'
import NumberField from './NumberField.vue'

function onClick(result: any) {
  console.log(`Clicked ${result}`);
}
</script>

<template>
  <table>
    <tbody>
      <tr v-for="row in 3" :key="row">
        <td v-for="col in 3" :key="col">
          <NumKeyItem :num="col + row * 3 - 3" :onClick />
        </td>
      </tr>
      <tr>
        <td>
          <CharKeyItem char="X" :onClick />
        </td>
      </tr>
    </tbody>
  </table>
</template>

<style lang="css" scoped>
div {
  background-color: red;
}

table {
  background-color: blue;
  transform: scale(2);
  transform-origin: center;
}
</style>