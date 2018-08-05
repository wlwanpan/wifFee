package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/gmap"
	"github.com/wlwanpan/wifFee/services/pb"
	"googlemaps.github.io/maps"
)

func GetCoffeeShops(w http.ResponseWriter, r *http.Request) (int, error) {
	params := mux.Vars(r)
	client := gmap.MapClient()

	parsedLocation := gmap.ParseLocation(params["latlng"])
	searchReq := &maps.TextSearchRequest{
		Query:    "cafe",
		Radius:   1,
		Location: parsedLocation,
	}

	resp, err := client.TextSearch(context.Background(), searchReq)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	places := []*pb.PlacePb{}
	for _, place := range resp.Results {
		places = append(places, mapsPlaceToPb(place))
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
	return http.StatusNoContent, errors.New("Not implemented yet...")
}

func mapsPlaceToPb(p maps.PlacesSearchResult) *pb.PlacePb {
	return &pb.PlacePb{
		ID:      p.ID,
		PlaceID: p.PlaceID,
		Name:    p.Name,
		Address: &pb.PlaceAddress{
			Lat:          p.Geometry.Location.Lat,
			Lng:          p.Geometry.Location.Lng,
			LocationType: p.Geometry.LocationType,
		},
	}
}
