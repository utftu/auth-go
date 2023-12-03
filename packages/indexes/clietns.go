package indexes

import (
	"context"
	"site/src/libs/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClientIndexes() {
	client := mongodb.Connect()
	collection := client.Database("auth").Collection("clients")

	collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: options.Index().SetUnique(true),
	})
}
