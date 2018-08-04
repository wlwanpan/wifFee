package routers

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wlwanpan/wifFee/services/middlewares"
)

func SetCORS(r *mux.Router) http.Handler {

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(r)
}

func InitRoutes() *mux.Router {

	router := mux.NewRouter()
	logger := middlewares.ErrorLogging

	router = SetPlacesRoutes(router, logger)

	return router
}
