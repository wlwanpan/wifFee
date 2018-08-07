package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
	mgo "gopkg.in/mgo.v2"
)

func SetWifiRoutes(r *mux.Router, db *mgo.Session) *mux.Router {

	m := middlewares.Middleware{}
	m.Chain(middlewares.Logger)
	m.SetDB(db)

	r.HandleFunc("/wifi/{latlng}", m.Then(m.UseDB(handlers.LogWifi))).Methods("POST")
	return r

}
