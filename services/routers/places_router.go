package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
)

func SetPlacesRoutes(r *mux.Router) *mux.Router {

	m := middlewares.Middleware{}
	m.Chain(middlewares.Logger)

	r.HandleFunc("/places/{latlng}", m.Then(handlers.GetCoffeeShops)).Methods("GET")
	r.HandleFunc("/place/{id}", m.Then(handlers.GetCoffeeShop)).Methods("GET")

	return r
}
