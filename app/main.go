package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qbem-repos/golang-arch-sample/interactor"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", interactor.PostPerson).Methods(http.MethodPost)
	router.HandleFunc("/", interactor.GetPerson).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}
