<template>
  <div id='map-drawer'
    v-bind:class="[revealDrawer ? 'reveal-drawer': '']"
    @click='revealDrawer = !revealDrawer'>

  </div>
</template>

<script>
export default {
  name: 'map-drawer',
  data() {
    return {
      revealDrawer: false
    }
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