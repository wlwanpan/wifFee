<template>
  <div id='google-map'>
    <GmapMap
      ref='gMap'
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
  data() {
    return {
      zoom: 15,
      mapType: 'roadmap'
    }
  },
  computed: {
    ...mapGetters({
      center: 'navigator/getCurrentCoords',
      markers: 'navigator/getMarkers',
      places: 'navigator/getPlaces'
    })
  },
  async mounted() {
    try {
      await this.$store.dispatch('navigator/updateCurrentCoords')
    }
    catch (err) {
      window.alert(err)
    }
  }
}
</script>

<style>
#google-map {
  height: 100%; width: 100%;
}
</style>