package gmap

import (
	"context"
	"errors"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

// Retrieve google-map API key from environment variable
func mapsApiKey() maps.ClientOption {

	apiKey := os.Getenv("GOOGLE_MAP_API_KEY")
	return maps.WithAPIKey(apiKey)
}

// Convert location string "lat,lng" to google maps location
func ParseLocation(location string) *maps.LatLng {
	l, err := maps.ParseLatLng(location)
	if err != nil {
		log.Println("Could not parse latitude,longitude")
	}
	return &l
}

// Return google map client < to move to middleware
func MapClient() *maps.Client {

	mapsApiKey := mapsApiKey()
	client, err := maps.NewClient(mapsApiKey)
	if err != nil {
		log.Println("Could not load google map client")
	}
	return client
}

// Wrapper to maps.TextSearch <- Need to update params to config struct to support more options.
func TextSearch(latlng string, radius uint) (maps.PlacesSearchResponse, error) {

	l := ParseLocation(latlng)
	ctx := context.Background()
	r := &maps.TextSearchRequest{
		Query:    "cafe",
		Radius:   radius,
		Location: l,
	}

	return MapClient().TextSearch(ctx, r)
}

// Wrapper to maps.PlaceDetails, takes in PlaceID
func PlaceDetails(id string) (maps.PlaceDetailsResult, error) {

	if id == "" {
		return maps.PlaceDetailsResult{}, errors.New("Invalid place ID")
	}
	r := &maps.PlaceDetailsRequest{PlaceID: id}
	ctx := context.Background()

	return MapClient().PlaceDetails(ctx, r)
}
