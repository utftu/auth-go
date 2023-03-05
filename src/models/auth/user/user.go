package user

import (
	"auth-go/src/models/auth/responses"
	"encoding/json"
)

type User struct {
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Avatar           string `json:"avatar" bson:"avatar"`
	OriginalResponse string `json:"original_response" bson:"original_response"`
}

func CreateUserFromProvider(provider string, data interface{}) *User {
	switch provider {
	case "do":
		{
			response := (data).(*responses.Do)
			return CreateUserFromDo(response)
		}
	}
	return nil
}

func CreateUserFromDo(do *responses.Do) *User {
	originalResponseJson, _ := json.Marshal(do)
	return &User{
		Name:             do.Info.Name,
		Email:            do.Info.Email,
		OriginalResponse: string(originalResponseJson),
	}
}

func CreateUserFromGoogle(do *responses.Do) *User {
	originalResponseJson, _ := json.Marshal(do)
	return &User{
		Name:             do.Info.Name,
		Email:            do.Info.Email,
		OriginalResponse: string(originalResponseJson),
	}
}
