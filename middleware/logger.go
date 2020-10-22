package middleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// Logger is a middleware that log every endpoint request
func Logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}
