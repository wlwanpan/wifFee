<template lang="pug">
  div#map-page.full-page
    google-map
    map-drawer
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