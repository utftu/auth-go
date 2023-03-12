package authGoCore

import (
	"auth-go-core/user"
)

type User = user.User

type StrategyData struct {
	ClientId           string
	ClientSecret       string
	RedirectUrl        string
	ServiceRedirectUrl string
}

type Strategy interface {
	GetUserRedirectUrl(data *StrategyData) string
	GetUserData(data *StrategyData, code string) (*user.User, error)
}

type AuthGoCore struct {
	Data     StrategyData
	Strategy Strategy
}

func (authGoCore *AuthGoCore) GetUserRedirectUrl() string {
	return authGoCore.Strategy.GetUserRedirectUrl(&authGoCore.Data)
}

func (authGoCore *AuthGoCore) GetUserData(code string) (*user.User, error) {
	return authGoCore.Strategy.GetUserData(&authGoCore.Data, code)
}
