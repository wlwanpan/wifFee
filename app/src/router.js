import Vue from 'vue'
import VueRouter from 'vue-router'
import MapPage from './pages/MapPage'

Vue.use(VueRouter)

export default new VueRouter({
  routes: [
    {
      path: '/', component: MapPage
    }
  ]
})