
export default {
  namespaced: true,
  state: {
    currentCoords: {
      lat: 49.281209,
      lng: -123.118336
    },
    mapType: 'roadmap',
    zoom: 15,
    markers: [],
    places: []
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
    }
  },
  actions: {
    updateCurrentCoords({commit}) {
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
    setCurrentCoords(state, coords) {
      console.log(coords)
      state.currentCoords = {
        lat: coords.latitude,
        lng: coords.longitude
      }
    }
  }
}