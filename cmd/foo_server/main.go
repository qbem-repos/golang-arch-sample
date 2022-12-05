package main

import (
	"log"

	"github.com/qbem-repository/golang-arch-sample/pkg/webapi"
)

func main() {
	// Todo: Pendings
	// worker.Run()

	if err := webapi.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
