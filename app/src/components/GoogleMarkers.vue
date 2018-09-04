<template lang="pug">
  div#google-markers
    GmapMarker(
      v-for="(m, index) in markers",
      :key="index",
      :position="m.position",
      :icon="m.icon",
      :clickable="true",
      :draggable="false",
      @mouseover="onHover(m.placeId)")
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'google-markers',
  computed: {
    ...mapGetters({
      center: 'navigator/getCurrentCoords',
      markers: 'marker/getMarkers'
    })
  },
  watch: {
    center() {
      this.updateMarkers()
    }
  },
  methods: {
    async onHover(placeId) {
      let markerDetails = await this.$store.dispatch('marker/loadDetails', placeId)
      console.log(markerDetails)
    },
    updateMarkers() {
      this.$store.dispatch('marker/updateMarkers', this.center)
    }
  }
}
</script>