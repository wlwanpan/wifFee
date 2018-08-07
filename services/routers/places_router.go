package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
	mgo "gopkg.in/mgo.v2"
)

func SetPlacesRoutes(r *mux.Router, db *mgo.Session) *mux.Router {

	m := middlewares.Middleware{}
	m.Chain(middlewares.Logger)
	m.SetDB(db)

	r.HandleFunc("/places/{latlng}", m.Then(m.UseDB(handlers.GetCoffeeShops))).Methods("GET")
	r.HandleFunc("/place/{id}", m.Then(m.UseDB(handlers.GetCoffeeShop))).Methods("GET")

	return r
}
