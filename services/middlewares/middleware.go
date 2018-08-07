package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Application handler returning status code and error
type AppHandler func(w http.ResponseWriter, r *http.Request) (int, error)

// Middleware function interface
type MiddlewareFunc func(h http.HandlerFunc) http.HandlerFunc

// Middleware used by router to set func chain for http handlers.
type Middleware struct {
	// Store an arr of middleware func to call on each handler func.
	chain []MiddlewareFunc
	// Holds a mongoDB session for handlers using db instances.
	db *mgo.Session
}

// Append middleware funcs to chain.
func (m *Middleware) Chain(fs ...MiddlewareFunc) {

	for _, f := range fs {
		m.chain = append(m.chain, f)
	}
}

// Loop through middleware chain and call Recoverer for error handling as last step.
func (m *Middleware) Then(f AppHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		h := Recoverer(f)

		for _, next := range m.chain {
			h = next(h)
		}

		h(w, r)
	}
}

// Stores a mongoDB session to middlware
func (m *Middleware) SetDB(db *mgo.Session) {
	m.db = db
}

// Make a copy of mongoDB stored in the middleware and store the Session
// in the request context.
func (m *Middleware) UseDB(h AppHandler) AppHandler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		dbSession := m.db.Copy()
		defer dbSession.Close()

		ctx := context.WithValue(r.Context(), "db", dbSession)
		r = r.WithContext(ctx)

		return h(w, r)
	}
}

// Logs incoming request, format: [METHOD], URL, DurationOfRequest
func Logger(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			log.Printf("[%s] %q %v", r.Method, r.URL.String(), time.Since(startTime))
		}()

		h(w, r)
	}
}

// Catches Application error and parse as http Error
func Recoverer(h AppHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		status, err := h(w, r)

		if err != nil {
			errMsg := err.Error()

			log.Println(errMsg)
			http.Error(w, errMsg, status)
		}
	}
}
