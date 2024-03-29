package client

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewClientMongo(client *mongo.Client) *ClientMongo {
	return &ClientMongo{
		client,
	}
}

type ClientMongo struct {
	client *mongo.Client
}

func (clientMongo *ClientMongo) GetCollection() *mongo.Collection {
	collection := clientMongo.client.Database("auth").Collection("clients")
	return collection
}

func (clientMongo *ClientMongo) Get(key string, id string) *Client {
	var client Client
	err := clientMongo.GetCollection().FindOne(context.Background(), bson.D{{
		Key: key, Value: id,
	}}).Decode(&client)
	if err != nil {
		return nil
	}

	return &client
}

func (clientMongo *ClientMongo) GetByName(name string) *Client {
	return clientMongo.Get("name", name)
}
