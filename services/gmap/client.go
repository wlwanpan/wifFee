package gmap

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

// Retrieve google-map API key from environment variable
func mapsApiKey() maps.ClientOption {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
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
