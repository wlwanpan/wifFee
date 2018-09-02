<template lang="pug">
  button(class="wif-btn", @click="onClick()") {{ label }}
</template>

<script>
export default {
  name: 'wif-button',
  props: {
    label: String,
    onClick: Function
  }
}
</script>

<style lang="scss" scoped>
.wif-btn {
  width: 80%;
  border-radius: 4px;
  background-color: white;
}
</style>