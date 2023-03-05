package utils

import (
	"auth-go/auth-go-core/user"
	"io"
	"net/http"
	"net/url"
)

type UserTokenData struct {
	Url string

	ClientId         string
	ClientSecret     string
	GrantType        string
	Code             string
	UserTokenUrl string
}

type handleToken func(io.Reader) (*user.User, error)

func GetUserData(dataToken *UserTokenData, handleToken handleToken) (*user.User, error) {
	parsedUrl, err := url.Parse(dataToken.Url)
	if err != nil {
		return nil, err
	}

	params := parsedUrl.Query()
	params.Add("client_id", dataToken.ClientId)
	params.Add("client_secret", dataToken.ClientSecret)
	params.Add("grant_type", dataToken.GrantType)
	params.Add("code", dataToken.Code)
	params.Add("redirect_uri", dataToken.UserTokenUrl)
	parsedUrl.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", parsedUrl.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	user, err := handleToken(res.Body)

	if err != nil {
		return nil, err
	}

	return user, nil
}
