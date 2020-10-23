package middleware

import (
	"net/http"

	"github.com/goji/httpauth"
)

// BasicAuthen is a middleware that
func BasicAuthen(name, password string, next http.Handler) http.Handler {
	return httpauth.SimpleBasicAuth(name, password)(next)
}
