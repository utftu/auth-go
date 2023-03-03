package auth

import (
	"auth-go/src/models/auth/responses"
	"auth-go/src/models/auth/user"
	"fmt"
	"net/http"

	"encoding/json"
	"io"
)                                                          

func HandleGoogleResponse(response io.Reader) (*user.User, error) {
	var googleRequest1 responses.GoogleRequest1

	err := json.NewDecoder(response).Decode(&googleRequest1)

	if err != nil {
		return nil, err
	}

	res, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", googleRequest1.AccessToken))

	if err != nil {
		return nil, err
	}

	var googleRequest2 responses.GoogleRequest2
	err = json.NewDecoder(res.Body).Decode(&googleRequest2)

	if err != nil {
		return nil, err
	}

	originalResponseJson, _ := json.Marshal(googleRequest2)
	return &user.User{
		Name: googleRequest2.Name,
		Email: googleRequest2.Email,
		Avatar: googleRequest2.Picture,
		OriginalResponse: string(originalResponseJson),
	}, nil
}