import { gmapApi } from 'vue2-google-maps'
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
      let res = await services().get('places/' + [lat, lng].join(','))
      if (res.statusText === 'OK') {
        commit('setPlaces', res.data.places)
      }
    },
    updateCurrentCoords({ commit }) {
      /**
       * Locate current posn using window.navigator and returns a promise{true} if successful
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
      let indexedPlaces = {}
      places.forEach((place) => {
        let place_id = place.place_id
        delete place.place_id
        indexedPlaces[place_id] = place
      })
      console.log(indexedPlaces)
      state.places = indexedPlaces
    },
    setCurrentCoords(state, coords) {
      state.currentCoords = {
        lat: coords.latitude,
        lng: coords.longitude
      }
    }
  }
}