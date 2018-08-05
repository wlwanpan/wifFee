import { gmapApi } from 'vue2-google-maps'
import pbParser from '../../pb/parser'
import services from '../../services'

export default {
  namespaced: true,
  state: {
    map: null,
    currentCoords: {
      lat: 49.281209,
      lng: -123.118336
    },
    mapType: 'roadmap',
    zoom: 15,
    markers: [],
    places: {}
  },
  getters: {
    getCurrentCoords(state) {
      return state.currentCoords
    },
    getMarkers(state) {
      return state.markers
    },
    getPlaces(state) {
      return state.places
    },
    getMapType(state) {
      return state.mapType
    },
    getZoom(state) {
      return state.zoom
    }
  },
  actions: {
    setMapInstance({ commit }, gmapInstance) {
      /**
       * Store google Map obj and initialize/store google service obj
       * @param {google.maps.Map}
       */
      commit('setMapInstance', gmapInstance)
    },
    async updateCoffeeShopMarkers({ commit, state }) {
      /**
       * Call Services search query
       */
      let {lat, lng} = state.currentCoords
      let resp = await services().get('places/' + [lat, lng].join(','))

      if (resp.status === 200) {
        commit('setPlaces', pbParser.placesToIndexedJson(resp.data))
      }
      else {
        console.log(resp)
      }
    },
    updateCurrentCoords({ commit }) {
      /**
       * Locate current posn using window.navigator and returns a promise{true} if successful
       * @return {Promise(boolean)}
       */
      return new Promise((resolve, reject) => {
        navigator.geolocation.getCurrentPosition(posn => {
          if (posn) {
            commit('setCurrentCoords', posn.coords)
            resolve(true)
          }
          else {
            reject(new Error('Could not find current location.'))
          }
        })
      })
    }
  },
  mutations: {
    setMapInstance(state, gmapInstance) {
      state.map = gmapInstance
    },
    setPlaces(state, places) {
      state.places = places
    },
    setCurrentCoords(state, coords) {
      state.currentCoords = {
        lat: coords.latitude,
        lng: coords.longitude
      }
    }
  }
}