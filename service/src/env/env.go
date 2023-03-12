package env

import "go.mongodb.org/mongo-driver/mongo"

type Env struct {
	Mongo *mongo.Client
}
