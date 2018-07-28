package main

import (
	"log"
	"net/http"

	"github.com/wlwanpan/wifFee/services/routers"
)

func main() {

	router := routers.InitRoutes()

	log.Fatal(http.ListenAndServe(":3000", router))

}
