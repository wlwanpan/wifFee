package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"googlemaps.github.io/maps"
)

type PlacesResponse struct {
	Places []maps.PlacesSearchResult `json:"places"`
}

func parseLocation(location string) *maps.LatLng {
	l, err := maps.ParseLatLng(location)
	if err != nil {
		log.Println("Could not parse latitude,longitude")
	}
	return &l
}

func GetCoffeeShops(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	client := MapClient()

	parsedLocation := parseLocation(params["latlng"])
	searchReq := &maps.TextSearchRequest{
		Query:    "cafe",
		Radius:   50,
		Location: parsedLocation,
	}

	resp, err := client.TextSearch(context.Background(), searchReq)
	if err != nil {
		log.Println(err)
	}

	// -> Change to protobuf
	placesJson, err := json.Marshal(PlacesResponse{Places: resp.Results})
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(placesJson)
}
