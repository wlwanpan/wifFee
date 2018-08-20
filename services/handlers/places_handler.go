package handlers

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/gmap"
	"github.com/wlwanpan/wifFee/services/models"
	"github.com/wlwanpan/wifFee/services/pb"
	"googlemaps.github.io/maps"
)

func GetCoffeeShops(w http.ResponseWriter, r *http.Request) (int, error) {
	params := mux.Vars(r)

	resp, err := gmap.TextSearch(params["latlng"], 10)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	places := []*pb.PlacePb{}
	for _, mapsPlace := range resp.Results {
		aLog := models.ActionLog{
			RecordID:   mapsPlace.PlaceID,
			Collection: "places",
			Action:     "load",
		}
		aLog.Create()

		place := models.Place{PlaceID: mapsPlace.PlaceID}
		err := place.Load()
		if err != nil {
			log.Println(err.Error())
		}

		places = append(places, PlaceToPb(place, mapsPlace))
	}

	placesPb := &pb.PlacesPb{
		Places: places,
	}
	encodedPlacesPb, err := proto.Marshal(placesPb)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(encodedPlacesPb)

	return http.StatusOK, nil
}

func GetCoffeeShop(w http.ResponseWriter, r *http.Request) (int, error) {
	placeID := mux.Vars(r)["id"]

	errChan := make(chan error)
	done := make(chan bool)

	var mapsPlace maps.PlaceDetailsResult
	var place models.Place

	go func() {
		p := models.Place{PlaceID: placeID}
		err := p.LoadOrCreate()
		if err != nil {
			errChan <- err
		}
		place = p
		done <- true
	}()

	go func() {
		resp, err := gmap.PlaceDetails(placeID)
		if err != nil {
			errChan <- err
		}
		mapsPlace = resp
		done <- true
	}()

	for i := 0; i < 2; i++ {
		select {
		case err := <-errChan:
			return http.StatusInternalServerError, err
		default:
			<-done
		}
	}

	encodedPlacePb, err := proto.Marshal(PlaceDetailsToPb(place, mapsPlace))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(encodedPlacePb)

	return http.StatusOK, nil
}

func PlaceDetailsToPb(place models.Place, mapsPlaceDetail maps.PlaceDetailsResult) *pb.PlacePb {
	return &pb.PlacePb{
		PlaceID: place.PlaceID,
	}
}

func PlaceToPb(place models.Place, mapsPlace maps.PlacesSearchResult) *pb.PlacePb {
	return &pb.PlacePb{
		PlaceID: place.PlaceID,
		Name:    mapsPlace.Name,
		Wifi: &pb.PlaceWifi{
			AvgUp:   place.AvgUp,
			AvgDown: place.AvgDown,
		},
		Address: &pb.PlaceAddress{
			Lat:          mapsPlace.Geometry.Location.Lat,
			Lng:          mapsPlace.Geometry.Location.Lng,
			LocationType: mapsPlace.Geometry.LocationType,
		},
	}
}
