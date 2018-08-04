package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/wlwanpan/wifFee/services/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables, .env file")
	}

	router := routers.InitRoutes()
	port := ":" + os.Getenv("PORT")

	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(port, routers.SetCORS(router)))

}
