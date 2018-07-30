<template>
  <div id='google-map'>
    <GmapMap
      ref='gMapRef'
      :center='center'
      :zoom='zoom'
      :map-type-id='mapType'
      style="width: 100%; height: 100%;"
    >
      <GmapMarker
        :key="index"
        v-for="(m, index) in markers"
        :position="m.position"
        :clickable="true"
        :draggable="true"
        @click="center=m.position"
      />
    </GmapMap>
  </div>
</template>

<script>
import { gmapApi } from 'vue2-google-maps'
import { mapGetters } from 'vuex'

export default {
  name: 'google-map',
  computed: {
    ...mapGetters({
      center: 'navigator/getCurrentCoords',
      markers: 'navigator/getMarkers',
      places: 'navigator/getPlaces',
      mapType: 'navigator/getMapType',
      zoom: 'navigator/getZoom'
    })
  },
  async mounted() {
    try {
      let map = await this.$refs.gMapRef.$mapPromise
      this.$store.dispatch('navigator/setMapInstance', map)
      await this.initMapSetup()
    }
    catch (err) {
      window.alert(err)
    }
  },
  methods: {
    async initMapSetup() {
      await this.$store.dispatch('navigator/updateCurrentCoords')
      await this.$store.dispatch('navigator/updateCoffeeShopMarkers')
    }
  }
}
</script>

<style>
#google-map {
  height: 100%; width: 100%;
}
</style>