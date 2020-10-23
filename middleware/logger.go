package middleware

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// Logger is a middleware that log every endpoint request
func Logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

// HandlerLogger is a middleware that log every endpoint request
func HandlerLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		url := r.URL.Path
		log.Printf("[%v] %v\n", method, url)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
