package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
	mgo "gopkg.in/mgo.v2"
)

func SetPlacesRoutes(r *mux.Router, c middlewares.Handler, db *mgo.Session) *mux.Router {

	r.HandleFunc("/places/{latlng}", c(handlers.GetCoffeeShops)).Methods("GET")
	r.HandleFunc("/place/{:id}", c(handlers.GetCoffeeShop)).Methods("GET")

	return r
}
