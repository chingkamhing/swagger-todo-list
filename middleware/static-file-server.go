package middleware

import (
	"net/http"
	"os"
	"path"
	"strings"
)

// API request and prefix; any GET request url under this path will be handlered with default API router
const apiRequest = "GET"
const apiPrefix = "/api"

// Serves index.html in case the requested file isn't found (or some other os.Stat error)
func serveIndex(assetPath string, serve http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage := path.Join(assetPath, "index.html")
		requestedPage := path.Join(assetPath, r.URL.Path)
		_, err := os.Stat(requestedPage)
		if err != nil {
			// serve index if page doesn't exist
			http.ServeFile(w, r, indexPage)
			return
		}
		serve.ServeHTTP(w, r)
	}
}

// StaticFileServer is a middleware that serve static file in public directory
func StaticFileServer(assetPath string, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		url := r.URL.Path
		if strings.HasPrefix(url, apiPrefix) {
			next.ServeHTTP(w, r)
		} else if method == apiRequest {
			serveIndex(assetPath, http.FileServer(http.Dir(assetPath))).ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(fn)
}
