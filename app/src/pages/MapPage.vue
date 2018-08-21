<template>
  <div id='map-page' class='full-page'>
    <google-map />
    <map-drawer />
  </div>
</template>

<script>
import GoogleMap from '../components/GoogleMap'
import MapDrawer from '../components/MapDrawer'

export default {
  components: {
    GoogleMap,
    MapDrawer
  }
}
</script>

<style lang='scss' scoped>
#map-page {
  position: relative;
}
</style>