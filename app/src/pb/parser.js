import { PlacesPb } from './places_pb'

export default {

  placesToIndexedJson(bytes) {
    /**
     * Convert places protobuf obj to json and index by place_id
     * @return {Object(placeid, {})}
     */
    let indexedPlaces = {}
    let places = []

    try {
      let placesPb = PlacesPb.deserializeBinary(bytes)
      places = placesPb.getPlacesList()
    }
    catch (e) {
      window.alert(e)
    }

    for (let i = 0; i < places.length; i ++) {
      let place = places[i].toObject()
      let _placeId = place.placeid
      delete place.placeid
      indexedPlaces[_placeId] = place
    }

    console.log(indexedPlaces)
    return indexedPlaces
  }

}