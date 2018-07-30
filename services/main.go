package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router)))

}
