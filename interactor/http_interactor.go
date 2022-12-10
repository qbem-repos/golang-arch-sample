package interactor

import (
	"net/http"

	"github.com/qbem-repos/golang-arch-sample/adapter"
	"github.com/qbem-repos/golang-arch-sample/provider"
	"github.com/qbem-repos/golang-arch-sample/service"
)

func PostPerson(w http.ResponseWriter, r *http.Request) {
	person := adapter.RequestToPersonDBDatamodel(r)
	service.Validation(person, w)
	provider.SavePerson(person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	provider.FindPersonsByFilter(adapter.QueryStringToMap(r.URL.Query()))
}
