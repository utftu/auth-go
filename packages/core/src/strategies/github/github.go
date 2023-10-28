package github

import (
	"core"
	"core/src/strategies/utils"
	"core/src/user"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Strategy struct {
}

func (doStrategy *Strategy) GetUserRedirectUrl(data *core.StrategyData) string {
	return utils.GetUserRedirectUrl(&utils.UserRedirectData{
		Url:          "https://github.com/login/oauth/authorize",
		ClientId:     data.ClientId,
		ResponseType: "code",
		RedirectUrl:  data.RedirectUrl,
		ServiceRedirectUrl: data.ServiceRedirectUrl,
		Scope:        []string{"read:user", "user:email"},
	})
}

func HandleResponse(res_body io.Reader) (*user.User, error) {
	var token Token
	json.NewDecoder(res_body).Decode(&token)

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
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

	req, err = http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client = &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	var emails []Email
	json.NewDecoder(res.Body).Decode(&emails)

	return &user.User{
		Name:   account.Login,
		Avatar: account.AvatarURL,
		Email:  emails[0].Email,
	}, nil
}

func (doStrategy *Strategy) GetUserData(data *core.StrategyData, code string) (*user.User, error) {
	return utils.GetUserData(&utils.UserTokenData{
		Url:          "https://github.com/login/oauth/access_token",
		ClientId:     data.ClientId,
		ClientSecret: data.ClientSecret,
		ServiceHandleUrl: data.ServiceRedirectUrl,
		GrantType:    "authorization_code",
		Code:         code,
	}, HandleResponse)
}
