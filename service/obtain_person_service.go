package service

import (
	"github.com/qbem-repos/golang-arch-sample/datastruct"
	"github.com/qbem-repos/golang-arch-sample/gateway/db"

	"go.mongodb.org/mongo-driver/bson"
)

func ObtainPersons() ([]datastruct.Person, error) {
	var persons []datastruct.Person
	col := db.Db().Collection("persons")
	cur, err := col.Find(db.DbContext(), bson.M{})

	if err != nil {
		return nil, err
	}

	for cur.Next(db.DbContext()) {
		var elem datastruct.Person
		err := cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		persons = append(persons, elem)
	}

	return persons, nil
}
