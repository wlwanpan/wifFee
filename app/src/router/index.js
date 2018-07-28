import Vue from 'vue'
import Router from 'vue-router'
import MapPage from '../components/MapPage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/', component: MapPage
    }
  ]
})