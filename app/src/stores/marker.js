import { gmapApi } from 'vue2-google-maps'
import services from '../services'
import pbParser from '../pb/parser'

export default {
  namespaced: true,
  state: {
    markers: {}
  },
  getters: {
    getMarkers: (state) => {
      var markers = []
      if (!gmapApi()) return markers

      for (let key in state.markers) {
        let place = state.markers[key]
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
    updateMarkers: async ({ commit }, { lat, lng }) => {
      let resp = await services().get('places/' + [lat, lng].join(','))
      if (resp.status < 400) {
        commit('setMarkers', pbParser.placesToIndexedJson(resp.data))
      }
      else {
        console.log(resp)
      }
    },
    loadDetails: async (context, placeId) => {
      let resp = await services().get('place/' + placeId)
      if (resp.status < 400) {
        return pbParser.placeDetails(resp.data)
      }
      else {
        window.alert(resp.data)
      }
    }
  },
  mutations: {
    setMarkers(state, markers) {
      state.markers = markers
    }
  }
}