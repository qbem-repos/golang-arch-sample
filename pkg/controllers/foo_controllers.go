package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qbem-repository/golang-arch-sample/internal/foo/entity"
	"github.com/qbem-repository/golang-arch-sample/internal/foo/repository"
	"github.com/qbem-repository/golang-arch-sample/internal/foo/usecase"
)

func GetFoo(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("foo"))
}

func PostFoo(w http.ResponseWriter, r *http.Request) {
	// MOCK
	jsonMock := `
	{
	"Foo": "value"
	}
	`
	var responseFoo entity.Foo
	repoFoo := new(repository.FooRepository)
	// END MOCK

	json.NewDecoder(strings.NewReader(jsonMock)).Decode(&responseFoo)

	usecase.NewSaveNewFooUsecase(*repoFoo)

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("{message: success}"))
}
