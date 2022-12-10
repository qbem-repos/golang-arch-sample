package service

import (
	"github.com/qbem-repos/golang-arch-sample/datastruct"
	"github.com/qbem-repos/golang-arch-sample/gateway/db"
)

func SavePerson(person datastruct.Person) error {
	col := db.Db().Collection("persons")
	_, err := col.InsertOne(db.DbContext(), person)
	return err
}
