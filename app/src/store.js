import Vue from 'vue'
import Vuex from 'vuex'
import navigator from './stores/navigator'
import marker from './stores/marker'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    navigator,
    marker
  }
})