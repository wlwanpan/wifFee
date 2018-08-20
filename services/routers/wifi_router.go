package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
)

func SetWifiRoutes(r *mux.Router) *mux.Router {

	m := middlewares.Middleware{}
	m.Chain(middlewares.Logger)

	r.HandleFunc("/wifi/{latlng}", m.Then(handlers.LogWifi)).Methods("POST")
	return r

}
