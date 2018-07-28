<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import GoogleMap from './components/GoogleMap'

export default {
  name: 'app',
  data () {
    return {}
  },
  components: {
    GoogleMap
  }
}
</script>

<style lang="scss">
@import '~normalize.css';

#app {
  position: absolute;
  top: 0; bottom: 0; left: 0; right: 0;
  width: 100%; height: 100%;
}
</style>
