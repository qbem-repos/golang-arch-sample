package routes

import (
	"net/http"

	"github.com/qbem-repository/golang-arch-sample/pkg/controllers"
)

var fooRoutes = []Route{
	{
		URI:          "/healthcheck",
		Method:       http.MethodGet,
		Function:     func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) },
		Autenticated: false,
	},
	{
		URI:          "/foo",
		Method:       http.MethodGet,
		Function:     controllers.GetFoo,
		Autenticated: false,
	},
	{
		URI:          "/foo",
		Method:       http.MethodPost,
		Function:     controllers.PostFoo,
		Autenticated: false,
	},
}
