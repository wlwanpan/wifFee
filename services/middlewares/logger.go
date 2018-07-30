package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	// Log handler request to console
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
	}

}
