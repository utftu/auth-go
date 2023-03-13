package userConnection

import (
	"context"
	"time"

	"service/src/libs/random"

	"auth-go-core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type UserMongo struct {
	User authGoCore.User `bson:",inline" json:",inline"`
	Code      string `json:"code" bson:"code"`
	Id        string `json:"id" bson:"_id"`
	Created primitive.DateTime `json:"created" bson:"created"`
}

type UserMongoInsert struct {
	authGoCore.User `bson:",inline" json:",inline"`
	Code      string `json:"code"`
	Created primitive.DateTime `json:"created" bson:"created"`
}

func NewUserMongoConnection(client *mongo.Client) *UserMongoConnection {
	return &UserMongoConnection{
		client,
	}
}

type UserMongoConnection struct {
	client *mongo.Client
}

func (clientMongo *UserMongoConnection) GetCollection() *mongo.Collection {
	collection := clientMongo.client.Database("auth").Collection("users")
	return collection
}

func (clientMongo *UserMongoConnection) Get(key string, id string) *UserMongo {
	var user UserMongo
	filter := bson.D{bson.E{Key: key, Value: id}}

	err := clientMongo.GetCollection().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil
	}

	return &user
}

func (clientMongo *UserMongoConnection) Save(user *authGoCore.User) string {
	code := random.GetRandString(50)

	collection := clientMongo.GetCollection()

	collection.InsertOne(context.Background(), UserMongoInsert{
		User: *user,
		Code: code,
		Created:  primitive.NewDateTimeFromTime(time.Now()),
	})

	return code
}
