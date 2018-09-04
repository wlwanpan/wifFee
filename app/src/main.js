import Vue from 'vue'
import * as VueGoogleMaps from 'vue2-google-maps'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'

import store from './store'
import router from './router'
import App from './App'

Vue.use(VueMaterial)
Vue.use(VueGoogleMaps, {
  load: {
    key: process.env.GOOGLE_MAP_API_KEY,
    libraries: 'places'
  }
})

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
