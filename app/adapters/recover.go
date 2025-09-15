package adapters

import (
	"log"
	"net/http"
)

// Recover from Panics
func Recover(debug bool) Adapter {
	return func(h http.Handler, response *interface{}) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Caught Panic: %+v", err)
					if debug {
						if str, ok := err.(string); ok {
							http.Error(w, str, 500)
						}
						return
					}
					http.Error(w, http.StatusText(500), 500)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}
