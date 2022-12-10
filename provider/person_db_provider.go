package provider

import (
	"context"
	"log"

	"github.com/qbem-repos/golang-arch-sample/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	err error
	ctx context.Context = context.TODO()

	client  *mongo.Client
	db      *mongo.Database
	persons *mongo.Collection
)

func init() {
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal()
	}

	db = client.Database("person_db")
	persons = db.Collection("persons")
}

// SavePerson save person in database
func SavePerson(person model.PersonDBModel) {
	persons.InsertOne(ctx, person)
}

// FindPersonsByFilter
func FindPersonsByFilter(filter map[string][]string) ([]model.PersonDBModel, error) {
	list := []model.PersonDBModel{}

	cur, err := persons.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.TryNext(ctx) {
		var person model.PersonDBModel
		err = cur.Decode(person)

		if err != nil {
			return nil, err
		}

		list = append(list, person)
	}

	return list, nil
}
