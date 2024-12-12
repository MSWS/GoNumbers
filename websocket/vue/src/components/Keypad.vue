<script setup lang="ts">
import { ref } from 'vue';
import CharKeyItem from './CharKeyItem.vue'
import NumberField from './NumberField.vue'

const props = defineProps<{
  onSubmit: (num: number) => void,
  max: number
}>()

const currentNum = ref("")

let { onSubmit, max } = props;

function onClick(char: string) {
  switch (char) {
    case "✓":
      onSubmit(Number(currentNum.value))
    case "X":
      break
    default:
      let wouldBe = currentNum.value + char

      if (Number(wouldBe) > max) {
        wouldBe = max.toString()
      }

      currentNum.value = wouldBe
      return
  }
  currentNum.value = ""
}
</script>

<template>
  <div>
    <NumberField>{{ currentNum }}</NumberField>
    <table>
      <tbody>
        <tr v-for="row in 3" :key="row">
          <td v-for="col in 3" :key="col">
            <CharKeyItem :char="(col + row * 3 - 3).toString()" :onClick />
          </td>
        </tr>
        <tr>
          <td v-for="char in ['X', '0', '✓']" :key="char">
            <CharKeyItem :char="char" :onClick />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style lang="css" scoped>
div {
  min-width: fit-content;
  min-height: fit-content;
}

table {
  width: 100%;
  height: 90%;
}

td {
  width: 33%;
  height: 25%;
}
</style>