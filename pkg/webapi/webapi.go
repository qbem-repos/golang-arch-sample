package webapi

import (
	"log"
	"net/http"

	"github.com/qbem-repository/golang-arch-sample/pkg/router"
)

func ListenAndServe() error {
	r := router.NewRouter()

	log.Println("Qbem WebServer [xxxxx] listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	return nil
}
