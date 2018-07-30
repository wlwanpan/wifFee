package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
)

func SetPlacesRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/places/{latlng}", middlewares.LogRequest(handlers.GetCoffeeShops)).Methods("GET")
	return router
}
