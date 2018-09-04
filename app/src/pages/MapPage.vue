<template lang="pug">
  div#map-page.full-page
    GmapMap(
      ref='gMapRef',
      :center='center',
      :zoom='zoom',
      :map-type-id='mapType',
      style="width: 100%; height: 100%;")
      google-markers
</template>

<script>
import * as _ from 'underscore'
import { mapGetters } from 'vuex'
import GoogleMarkers from '../components/GoogleMarkers'

export default {
  name: 'map-page',
  computed: {
    ...mapGetters({
      map: 'navigator/getMap',
      center: 'navigator/getCurrentCoords',
      mapType: 'navigator/getMapType',
      zoom: 'navigator/getZoom'
    })
  },
  async mounted() {
    try {
      let map = await this.$refs.gMapRef.$mapPromise
      this.$store.dispatch('navigator/setMapInstance', map)
      await this.$store.dispatch('navigator/updateCurrentLoc')
      this.initListeners()
    }
    catch (err) {
      window.alert(err)
    }
  },
  methods: {
    initListeners() {
      let debouncedCb = _.debounce(() => {
        this.$store.dispatch('navigator/updateCurrentCoords', {
          latitude: this.map.getCenter().lat(),
          longitude: this.map.getCenter().lng()
        })
      }, 250)

      this.map.addListener('center_changed', debouncedCb)
    }
  },
  components: {
    GoogleMarkers
  }
}
</script>

<style lang='scss' scoped>
</style>