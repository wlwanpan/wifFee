import { gmapApi } from 'vue2-google-maps'
import pbParser from '../pb/parser'
import services from '../services'

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
    places: {}
  },
  getters: {
    getMap: (state) => state.map,
    getCurrentCoords: (state) => state.currentCoords,
    getMapType: (state) => state.mapType,
    getZoom: (state) => state.zoom,
    getPlaces: (state) => state.places,
    getMarkers: (state) => {
      var markers = []
      if (!gmapApi()) return markers

      for (let key in state.places) {
        let place = state.places[key]
        markers.push({
          placeId: key,
          position: {
            lat: place.address.lat,
            lng: place.address.lng
          },
          icon: {
            path: gmapApi().maps.SymbolPath.CIRCLE,
            scale: 5
          }
        })
      }
      return markers
    }
  },
  actions: {
    setMapInstance: ({ commit }, gmapInstance) => {
      /**
       * Store google Map obj and initialize/store google service obj
       * @param {google.maps.Map}
       */
      commit('setMapInstance', gmapInstance)
    },
    updateCoffeeShopMarkers: async ({ commit, state }) => {
      /**
       * Call Services search query
       */
      let {lat, lng} = state.currentCoords
      let resp = await services().get('places/' + [lat, lng].join(','))

      if (resp.status < 400) {
        commit('setPlaces', pbParser.placesToIndexedJson(resp.data))
      }
      else {
        console.log(resp)
      }
    },
    updateCurrentLoc: ({ commit }) => {
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
    },
    updateCurrentCoords: ({ commit }, coords) => {
      commit('setCurrentCoords', coords)
    }
  },
  mutations: {
    setMapInstance: (state, gmapInstance) => {
      state.map = gmapInstance
    },
    setPlaces: (state, places) => {
      state.places = places
    },
    setCurrentCoords: (state, coords) => {
      state.currentCoords = {
        lat: coords.latitude,
        lng: coords.longitude
      }
    }
  }
}