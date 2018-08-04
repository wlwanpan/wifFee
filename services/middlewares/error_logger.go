package middlewares

import (
	"log"
	"net/http"
	"time"
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request) (int, error)

type Handler func(n ErrorHandler) http.HandlerFunc

func ErrorLogging(next ErrorHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		defer func() {
			log.Printf("[%s] %q %v", r.Method, r.URL.String(), time.Since(startTime))
		}()

		status, err := next(w, r)

		if err != nil {
			http.Error(w, err.Error(), status)
		}
	}

}
