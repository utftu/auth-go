package main

import (
	"auth-go/src/libs/mongodb"
	"auth-go/src/models/auth/user/connection"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Student struct {
	FirstName string          `bson:"first_name,omitempty"`
	LastName  string          `bson:"last_name,omitempty"`
	Address   mongodb.Address `bson:"inline"`
	Age       int
}

func main() {
	client := mongodb.Connect()

	coll := client.Database("auth").Collection("users")
	// filter := bson.D{    bson.E{Key: "code", Value: "WM7CiFh05kfS3wywnniTsFumspZiKpE1e8nvyUJAUihp92K6CN"},}
	var result connection.UserMongo
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("-----", "err", err)
	fmt.Println(result)
}

// func main() {
// 	client := mongodb.Connect()

// 	coll := client.Database("school").Collection("students")
// 	filter := bson.D{{"age", 8}}
// 	var result Student
// 	err := coll.FindOne(context.TODO(), filter).Decode(&result)
// 	fmt.Println("-----", "err", err);
// 	fmt.Println(result.Address)
// }
