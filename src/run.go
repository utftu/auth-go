package main

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"auth-go/src/models/auth/user"
)

type UserMongo struct {
	user.User `bson:",inline" json:",inline"`
	Code      string `json:"code" bson:"code"`
	Id        string `json:"id" bson:"_id"`
}

func main() {
	// Simulate BSON data from MongoDB
	bsonData := primitive.M{
		"name":              "John Doe",
		"email":             "johndoe@example.com",
		"avatar":            "https://example.com/avatar.jpg",
		"original_response": `{"status":"success"}`,
		"code":              "1234",
		"_id":               "abc123",
	}

	// Convert the BSON data to a byte slice
	bsonBytes, err := bson.Marshal(bsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the BSON data to a UserMongo object
	var userMongo UserMongo
	err = bson.Unmarshal(bsonBytes, &userMongo)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the UserMongo object to JSON
	jsonData, err := json.Marshal(userMongo)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}
