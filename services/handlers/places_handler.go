package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/gmap"
	"googlemaps.github.io/maps"
)

type PlacesResponse struct {
	Places []maps.PlacesSearchResult `json:"places"`
}

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

	// -> Change to protobuf
	placesJson, err := json.Marshal(PlacesResponse{Places: resp.Results})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(placesJson)

	return http.StatusOK, nil
}
