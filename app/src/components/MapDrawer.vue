<template lang="pug">
  div#map-drawer(
    v-bind:class="[revealDrawer ? 'reveal-drawer': '']",
    @click="revealDrawer = !revealDrawer")
    upload-wifi
    search-place
</template>

<script>
import UploadWifi from './MapDrawer/UploadWifi'
import SearchPlace from './MapDrawer/SearchPlace'

export default {
  name: 'map-drawer',
  data() {
    return {
      revealDrawer: false
    }
  },
  components: {
    UploadWifi,
    SearchPlace
  }
}
</script>

<style lang='scss' scoped>
#map-drawer {
  position: fixed;
  index: 100; right: -300px;
  width: 320px; height: 100%;
  background-color: rgba(255, 255, 255, 1);

  &.reveal-drawer {
    right: 0px;
  }
}
</style>
