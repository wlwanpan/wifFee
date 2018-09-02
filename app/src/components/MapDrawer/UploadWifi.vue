<template lang="pug">
  div#upload-wifi Upload Wifi
    wif-button(
      :onClick="onClickHandler"
      label="Upload")
</template>

<script>
import { WifButton } from '../../elements'

export default {
  name: 'upload-wifi',
  data() {
    return {
    }
  },
  methods: {
    onClickHandler() {
      console.log('upload')
    }
  },
  components: {
    WifButton
  }
}
</script>

<style lang="scss" scoped>
#upload-wifi {

}
</style>