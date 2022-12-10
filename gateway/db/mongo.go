package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
)

func init() {
	ctx = context.Background()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	db = client.Database("persons_db")

}

func Db() *mongo.Database {
	return db
}

func DbContext() context.Context {
	return ctx
}
