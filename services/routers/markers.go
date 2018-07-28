package routers

import (
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/controllers"
)

func SetMarkerRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/markers", controllers.GetMarkers).Methods("GET")
	return router
}
