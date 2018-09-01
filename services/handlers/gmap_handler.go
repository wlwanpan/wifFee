package handlers

import (
	"context"
	"errors"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

var (
	gmapApiKey string = os.Getenv("GOOGLE_MAP_API_KEY")
)

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
	clientOption := maps.WithAPIKey(gmapApiKey)
	client, err := maps.NewClient(clientOption)
	if err != nil {
		log.Println("Could not load google map client")
	}

	return client
}

// Wrapper to maps.TextSearch <- Need to update params to config struct to support more options.
func TextSearch(latlng string, radius uint) ([]maps.PlacesSearchResult, error) {
	loc := ParseLocation(latlng)
	ctx := context.Background()
	req := &maps.TextSearchRequest{
		Query:    "cafe",
		Radius:   radius,
		Location: loc,
	}
	resp, err := MapClient().TextSearch(ctx, req)
	if err != nil {
		return []maps.PlacesSearchResult{}, err
	}

	return resp.Results, nil
}

// Wrapper to maps.PlaceDetails, takes in PlaceID
func PlaceDetails(id string) (maps.PlaceDetailsResult, error) {
	if id == "" {
		return maps.PlaceDetailsResult{}, errors.New("Invalid place ID")
	}
	req := &maps.PlaceDetailsRequest{PlaceID: id}
	ctx := context.Background()

	return MapClient().PlaceDetails(ctx, req)
}
