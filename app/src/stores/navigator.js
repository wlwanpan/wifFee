export default {
  namespaced: true,
  state: {
    map: null,
    currentCoords: {
      lat: 49.281209,
      lng: -123.118336
    },
    mapType: 'roadmap',
    zoom: 15
  },
  getters: {
    getMap: (state) => state.map,
    getCurrentCoords: (state) => state.currentCoords,
    getMapType: (state) => state.mapType,
    getZoom: (state) => state.zoom
  },
  actions: {
    setMapInstance: ({ commit }, gmapInstance) => {
      /**
       * Store google Map obj and initialize/store google service obj
       * @param {google.maps.Map}
       */
      commit('setMapInstance', gmapInstance)
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
    setCurrentCoords: (state, coords) => {
      state.currentCoords = {
        lat: coords.latitude,
        lng: coords.longitude
      }
    }
  }
}