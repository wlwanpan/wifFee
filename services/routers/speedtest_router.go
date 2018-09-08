package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/handlers"
	"github.com/wlwanpan/wifFee/services/middlewares"
)

func SetSpeedtestRoutes(r *mux.Router) *mux.Router {
	m := middlewares.Middleware{}
	m.Chain(middlewares.Logger)

	r.HandleFunc("/speedtest/ping", m.Then(handlers.Ping)).Methods("GET")
	r.HandleFunc("/speedtest/upload", m.Then(handlers.Upload)).Methods("POST")
	r.HandleFunc("/speedtest/download", m.Then(handlers.Download)).Methods("POST")
	return r
}
