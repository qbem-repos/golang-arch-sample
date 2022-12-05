package routes

import (
	"net/http"

	"github.com/qbem-repository/golang-arch-sample/pkg/controllers"
)

var barRoutes = []Route{
	{
		URI:          "/healthcheck",
		Method:       http.MethodGet,
		Function:     func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) },
		Autenticated: false,
	},
	{
		URI:          "/bar",
		Method:       http.MethodGet,
		Function:     controllers.GetBar,
		Autenticated: false,
	},
	{
		URI:          "/bar",
		Method:       http.MethodPost,
		Function:     controllers.PostBar,
		Autenticated: false,
	},
}
