package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qbem-repos/golang-arch-sample/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.PostPerson).Methods(http.MethodPost)
	router.HandleFunc("/", handler.GetPersons).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}
