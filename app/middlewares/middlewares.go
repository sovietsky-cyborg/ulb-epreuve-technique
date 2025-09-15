package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func RoutesMiddlewares(middlewares ...Middleware) func(h http.Handler)  http.Handler {
	return func(h http.Handler) http.Handler{

		for i := len(middlewares)-1; i >= 0; i--{
			h = middlewares[i](h)
		}

		return h
	}
}
