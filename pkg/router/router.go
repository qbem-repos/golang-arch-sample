package router

import (
	"github.com/gorilla/mux"
	"github.com/qbem-repository/golang-arch-sample/pkg/router/routes"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRouter(r)
}
