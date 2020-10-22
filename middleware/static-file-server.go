package middleware

import (
	"log"
	"net/http"
)

// StaticFileServer is a middleware that serve static file in public directory
func StaticFileServer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		uri := r.RequestURI
		log.Printf("[%s] %s\n", method, uri)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
