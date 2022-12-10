package service

import (
	"net/http"

	"github.com/go-playground/validator"
)

func validate(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	return err.(validator.ValidationErrors)
}

func Validation(s interface{}, w http.ResponseWriter) {
	if err := validate(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
