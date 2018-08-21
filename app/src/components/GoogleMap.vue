<template>
  <div id='google-map' class='full-page'>
    <GmapMap
      ref='gMapRef'
      :center='center'
      :zoom='zoom'
      :map-type-id='mapType'
      style="width: 100%; height: 100%;">
      <google-markers />
    </GmapMap>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import GoogleMarkers from './GoogleMarkers'

export default {
  name: 'google-map',
  computed: {
    ...mapGetters({
      center: 'navigator/getCurrentCoords',
      places: 'navigator/getPlaces',
      mapType: 'navigator/getMapType',
      zoom: 'navigator/getZoom'
    })
  },
  async mounted() {
    try {
      let map = await this.$refs.gMapRef.$mapPromise
      this.$store.dispatch('navigator/setMapInstance', map)
      await this.$store.dispatch('navigator/updateCurrentCoords')
    }
    catch (err) {
      window.alert(err)
    }
  },
  components: {
    GoogleMarkers
  }
}
</script>

<style lang='scss' scoped>
#google-map {
  position: absolute;
}
</style>