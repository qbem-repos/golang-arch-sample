package controllers

import "net/http"

func GetBar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
}

func PostBar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
	w.WriteHeader(http.StatusCreated)
}
