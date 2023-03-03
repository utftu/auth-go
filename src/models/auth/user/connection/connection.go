package connection

import (
	"context"
	"fmt"

	"auth-go/src/libs/random"
	"auth-go/src/models/auth/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	user.User `bson:",inline" json:",inline"`
	Code      string `json:"code" bson:"code"`
	Id        string `json:"id" bson:"_id"`
}

type UserMongoInsert struct {
	user.User  `bson:",inline" json:",inline"`
	Code string `json:"code"`
}

func NewUserMongoConnection(client *mongo.Client) *UserMongoConnection {
	return &UserMongoConnection {
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
	fmt.Println("-----", "key", key);
	fmt.Println("-----", "id", id);
	// var any interface {}
	// filter := bson.D{    bson.E{Key: key, Value: id},}
	// filter := bson.D{    bson.E{Key: key, Value: id},}
	filter := bson.D{    bson.E{Key: key, Value: id},}


	err := clientMongo.GetCollection().FindOne(context.TODO(), filter).Decode(&user)
  // err := clientMongo.GetCollection().FindOne(context.Background(), bson.D {{
	// 	Key: key, Value: id,
	// }}).Decode(&user)
	fmt.Println("-----", "user", user.User);
	// fmt.Println("-----", "any", any);
	if (err != nil) {
		return nil
	}

	return & user
}

func (clientMongo *UserMongoConnection) Save(user *user.User) string {
	code := random.GetRandString(50)

	collection := clientMongo.GetCollection()
	collection.InsertOne(context.Background(), UserMongoInsert {
		User: *user,
		Code: code,
	})

	return code
}