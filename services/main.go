package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/wlwanpan/wifFee/services/routers"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables, .env file")
	}

	db, err := mgo.Dial(os.Getenv("MONGO_DB_ADDRESS"))
	if err != nil {
		log.Fatal("Cannot connect to mongodb", err)
	}
	defer db.Close()

	port := ":" + os.Getenv("PORT")
	router := routers.Init(db)
	server := &http.Server{
		Addr:         port,
		Handler:      routers.SetCORS(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Starting server on port " + port)
	log.Fatal(server.ListenAndServe())

}
