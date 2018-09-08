<template lang="pug">
  div
    md-button.md-icon-button(
      @click="speedTest")
      md-icon cached
    md-button.md-icon-button(
      @click="navToAccount")
      md-icon account_circle
</template>

<script>
var speedTest = require('speedtest-net')

export default {
  name: 'drawer-right-actions',
  methods: {
    speedTest() {
      let test = speedTest({maxTime: 5000})

      test.on('data', (data) => {
        console.log(data)
      })
    },
    navToAccount() {

    }
  }
}
</script>