package auth

import (
	"auth-go/src/models/auth/responses"
	"auth-go/src/models/auth/user"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleGithubResponse(res_body io.Reader) (*user.User, error) {
	var tokenResponse responses.GithubToken
	json.NewDecoder(res_body).Decode(&tokenResponse)

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Set authorization header
	req.Header.Set("Accept", "application/json")

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenResponse.AccessToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var userGithub responses.GithubUser
	json.NewDecoder(res.Body).Decode(&userGithub)

	req, err = http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tokenResponse.AccessToken)

	client = &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	var emails []responses.GithubEmail
	json.NewDecoder(res.Body).Decode(&emails)

	fmt.Println("-----", "emails[0].Email", emails[0].Email)

	return &user.User{
		Name:   userGithub.Login,
		Avatar: userGithub.AvatarURL,
		Email:  emails[0].Email,
	}, nil
}
