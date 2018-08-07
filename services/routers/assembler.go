package routers

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func SetCORS(r *mux.Router) http.Handler {

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(r)
}

func Init(db *mgo.Session) *mux.Router {

	router := mux.NewRouter()
	router = SetPlacesRoutes(router, db)

	return router
}
