package do

import (
	authGoCore "auth-go-core"
	"auth-go-core/strategies/utils"
	"auth-go-core/user"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Strategy struct {
}

func (doStrategy *Strategy) GetUserRedirectUrl(data *authGoCore.StrategyData) string {
	return utils.GetUserRedirectUrl(&utils.UserRedirectData{
		Url:          "https://cloud.digitalocean.com/v1/oauth/authorize",
		ClientId:     data.ClientId,
		ResponseType: "code",
		RedirectUrl:  data.RedirectUrl,
		ServiceRedirectUrl: data.ServiceRedirectUrl,
		Scope:        []string{"read"},
	})
}

func HandleResponse(response io.Reader) (*user.User, error) {
	var token Token
	err := json.NewDecoder(response).Decode(&token)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://api.digitalocean.com/v2/account", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var account Account
	json.NewDecoder(res.Body).Decode(&account)

	req, err = http.NewRequest("GET", fmt.Sprintf("https://api.digitalocean.com/v2/users/%s", account.Account.UUID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	client = &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return &user.User{
		Name:             token.Info.Name,
		Email:            token.Info.Email,
	}, nil
}

func (doStrategy *Strategy) GetUserData(data *authGoCore.StrategyData, code string) (*user.User, error) {
	return utils.GetUserData(&utils.UserTokenData{
		Url:          "https://cloud.digitalocean.com/v1/oauth/token",
		ClientId:     data.ClientId,
		ClientSecret: data.ClientSecret,
		GrantType:    "authorization_code",
		ServiceHandleUrl: data.ServiceRedirectUrl,
		Code:         code,
	}, HandleResponse)
}
