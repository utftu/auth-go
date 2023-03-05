package google

import (
	authGoCore "auth-go/auth-go-core"
	"auth-go/auth-go-core/strategies/utils"
	"auth-go/auth-go-core/user"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Strategy struct {
}

func (doStrategy *Strategy) GetUserRedirectUrl(data *authGoCore.StrategyData) string {
	return utils.GetUserRedirectUrl(&utils.UserRedirectData{
		Url:          "https://accounts.google.com/o/oauth2/v2/auth",
		ClientId:     data.ClientId,
		ResponseType: "code",
		RedirectUrl:  data.RedirectUrl,
		UserRedirectUrl: data.UserRedirectUrl,
		Scope:        []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	})
}

func HandleResponse(response io.Reader) (*user.User, error) {
	var token Token

	err := json.NewDecoder(response).Decode(&token)

	if err != nil {
		return nil, err
	}

	res, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", token.AccessToken))

	if err != nil {
		return nil, err
	}

	var account Account
	err = json.NewDecoder(res.Body).Decode(&account)

	if err != nil {
		return nil, err
	}

	originalResponseJson, _ := json.Marshal(account)
	return &user.User{
		Name:             account.Name,
		Email:            account.Email,
		Avatar:           account.Picture,
		OriginalResponse: string(originalResponseJson),
	}, nil
}

func (doStrategy *Strategy) GetUserData(data *authGoCore.StrategyData, code string) (*user.User, error) {
	return utils.GetUserData(&utils.UserTokenData{
		Url:          "https://oauth2.googleapis.com/token",
		ClientId:     data.ClientId,
		ClientSecret: data.ClientSecret,
		GrantType:    "authorization_code",
		Code:         code,
	}, HandleResponse)
}
