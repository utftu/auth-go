package auth

import (
	"auth-go/src/models/auth/user"
	"fmt"
	"net/http"
	"net/url"
)

type Stage2 struct {
	Url  string
	Name string

	ClientId     string
	ClientSecret string
	GrantType    string
	Code         string
	RedirectUri  string
	State        string
}

func (callback *Stage2) Request() (*user.User, error) {
	parsedUrl, err := url.Parse(callback.Url)
	if err != nil {
		return nil, err
	}
	fmt.Println("-----", "callback.Code", callback.Code)
	params := parsedUrl.Query()
	params.Add("client_id", callback.ClientId)
	params.Add("client_secret", callback.ClientSecret)
	params.Add("grant_type", callback.GrantType)
	params.Add("code", callback.Code)
	params.Add("state", callback.State)
	params.Add("redirect_uri", callback.RedirectUri)
	parsedUrl.RawQuery = params.Encode()

	fmt.Println("-----", "parsedUrl.String()", parsedUrl.String())

	req, err := http.NewRequest("POST", parsedUrl.String(), nil)
	req.Header.Set("Content-Type", "application/json")

	// Set authorization header
	req.Header.Set("Accept", "application/json")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)

	// response, err := http.Post(parsedUrl.String(), "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var user *user.User
	switch callback.Name {
	case "do":
		{
			user, err = HandleDoResponse(res.Body)
		}
	case "google":
		{
			user, err = HandleGoogleResponse(res.Body)
		}
	case "github":
		{
			user, err = HandleGithubResponse(res.Body)
		}
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
