package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/wlwanpan/wifFee/services/models"
	"github.com/wlwanpan/wifFee/services/routers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables, .env file")
	}

	mgoAddr := os.Getenv("MONGO_DB_ADDRESS")
	models.InitDb(mgoAddr)

	port := ":" + os.Getenv("PORT")
	router := routers.AssembleRoutes()
	server := &http.Server{
		Addr:         port,
		Handler:      routers.SetCORS(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Starting server on port " + port)
	log.Fatal(server.ListenAndServe())

}
