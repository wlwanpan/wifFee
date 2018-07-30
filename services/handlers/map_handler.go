package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

func envGoogleApiKey() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("GOOGLE_MAP_API_KEY")
}

func MapClient() *maps.Client {

	envGoogleApiKey := envGoogleApiKey()
	mapsApiKey := maps.WithAPIKey(envGoogleApiKey)

	client, err := maps.NewClient(mapsApiKey)
	if err != nil {
		log.Println("Could not load google map client")
	}
	return client
}
