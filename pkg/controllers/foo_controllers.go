package controllers

import "net/http"

func GetFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

func PostFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
	w.WriteHeader(http.StatusCreated)
}
