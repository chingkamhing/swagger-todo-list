package middleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// Recover is a middleware that recover any panic endpoint
func Recover(next http.Handler) http.Handler {
	return middleware.Recoverer(next)
}
