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
      markers: 'navigator/getMarkers'
    })
  },
  watch: {
    center() {
      this.updateMarkers()
    }
  },
  methods: {
    onHover(placeId) {
      console.log(placeId)
    },
    updateMarkers() {
      this.$store.dispatch('navigator/updateCoffeeShopMarkers')
    }
  }
}
</script>