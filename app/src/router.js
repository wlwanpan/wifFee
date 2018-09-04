import Vue from 'vue'
import VueRouter from 'vue-router'

import MapPage from './pages/MapPage'
import AccountPage from './pages/AccountPage'
import ApiPage from './pages/ApiPage'
import AboutPage from './pages/AboutPage'

Vue.use(VueRouter)

export default new VueRouter({
  routes: [
    {
      path: '/', component: MapPage,
    },
    {
      path: '/account', component: AccountPage
    },
    {
      path: '/api', component: ApiPage
    },
    {
      path: '/about', component: AboutPage
    }
  ]
})