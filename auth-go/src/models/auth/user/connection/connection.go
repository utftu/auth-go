package userConnection

import (
	"context"

	"auth-go/src/libs/random"
	"github.com/utftufutukgyftryidytftuv/auth-go/auth-go-core"

	// "auth-go/src/models/auth/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	User authGoCore.User `bson:",inline" json:",inline"`
	Code      string `json:"code" bson:"code"`
	Id        string `json:"id" bson:"_id"`
}

type UserMongoInsert struct {
	authGoCore.User `bson:",inline" json:",inline"`
	Code      string `json:"code"`
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
	})

	return code
}
