package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ucl-epreuve-technique/app/utils"
)

func API(debug bool) Adapter {
	return func(h http.Handler, response *interface{}) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			h.ServeHTTP(w, r)
			w.Header().Set("Content-Type", "application/json")
			var payload = make(map[string]interface{})

			if e, ok := (*response).(utils.StatusError); ok {
				fmt.Printf("handler returned error code %v with message %v\n", e.Status(), e.Error()) // debug
				http.Error(w, e.Error(), e.Status())
			} else if s, ok := (*response).(fmt.Stringer); ok {
				payload["data"] = s.String()
				json.NewEncoder(w).Encode(payload)
			} else {
				payload["data"] = response
				json.NewEncoder(w).Encode(payload)
			}

		})
	}
}
