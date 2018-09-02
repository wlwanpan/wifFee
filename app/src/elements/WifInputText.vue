<template lang="pug">
  input.text-field(type="text", placeholder="placeholder")
</template>

<script>
export default {
  name: 'wif-input-text',
  props: {
    placeholder: String
  }
}
</script>

<style lang="scss">
.text-field {
  width: 80%;
}
</style>
