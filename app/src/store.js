import Vue from 'vue'
import Vuex from 'vuex'
import navigator from './stores/navigator'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    navigator
  }
})