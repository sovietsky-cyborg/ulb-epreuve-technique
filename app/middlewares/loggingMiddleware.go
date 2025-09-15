package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//// Logging all request to this endpoint
func Logging(l *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("Logging")

			l.Println("http:", r.Method, r.URL.Path, r.UserAgent())
			h.ServeHTTP(w, r)
		})
	}
}