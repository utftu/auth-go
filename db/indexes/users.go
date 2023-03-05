package indexes

import (
	"auth-go/src/libs/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUserIndexes() {
	client := mongodb.Connect()
	collection := client.Database("auth").Collection("users")

	collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"code": 1,
		},
		Options: options.Index().SetUnique(true),
	})
}
