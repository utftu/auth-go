package auth

import (
	"auth-go/src/models/auth/responses"
	"auth-go/src/models/auth/user"
	"fmt"
	"net/http"

	"encoding/json"
	"io"
)

func HandleDoResponse(response io.Reader) (*user.User, error) {
	var do responses.Do
	err := json.NewDecoder(response).Decode(&do)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://api.digitalocean.com/v2/account", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", do.AccessToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var doAccount responses.DoAccount
	json.NewDecoder(res.Body).Decode(&doAccount)

	req, err = http.NewRequest("GET", fmt.Sprintf("https://api.digitalocean.com/v2/users/%s", doAccount.Account.UUID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", do.AccessToken))

	client = &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, _ := io.ReadAll(res.Body)
	fmt.Println("-----", "bytes", string(bytes))

	originalResponseJson, _ := json.Marshal(do)
	return &user.User{
		Name:             do.Info.Name,
		Email:            do.Info.Email,
		OriginalResponse: string(originalResponseJson),
	}, nil
}
