package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Representa todas das routes da API
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	Autenticated bool
}

// Configurar todas as rotas do routes.
func ConfigRouter(r *mux.Router) *mux.Router {
	routes := fooRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
