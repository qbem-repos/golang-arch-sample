package handler

import (
	"encoding/json"
	"net/http"

	"github.com/qbem-repos/golang-arch-sample/service"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	res, err := service.ObtainPersons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(res)
}
