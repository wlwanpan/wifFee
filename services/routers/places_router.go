package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
)

func SetPlacesRoutes(r *mux.Router, c middlewares.Handler) *mux.Router {

	r.HandleFunc("/places/{latlng}", c(handlers.GetCoffeeShops)).Methods("GET")
	return r
}
