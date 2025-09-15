package adapters

import (
	"log"
	"net/http"
)

type Adapter func(http.Handler, *interface{}) (h http.Handler)

func Adapt(handler interface{}, adapters ...Adapter) (h http.Handler) {

	var response interface{}

	switch handler := handler.(type) {

	case http.Handler:
		h = handler
	case func(http.ResponseWriter, *http.Request):
		h = http.HandlerFunc(handler)
	case func(*http.Request) interface{}:
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response = handler(r)
		})
	case func(http.ResponseWriter, *http.Request) interface{}:
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response = handler(w, r)
		})
	default:
		log.Fatal("Invalid Adapt Handler", handler)
	}
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h, &response)
	}

	return h
}
