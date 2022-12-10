package adapter

import (
	"encoding/json"
	"net/http"

	datamodel "github.com/qbem-repos/golang-arch-sample/model"
)

func RequestToPersonDBDatamodel(r *http.Request) datamodel.PersonDBModel {
	person := new(datamodel.PersonDBModel)
	json.NewDecoder(r.Body).Decode(&person)
	return *person
}
