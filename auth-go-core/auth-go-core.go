package auth

type StrategyData struct {
	ClientId        string
	ClientSecret    string
	RedirectUrl     string
	UserRedirectUrl string
	UserTokenUrl    string
}

type Strategy interface {
	GetUserRedirectUrl(data *StrategyData) string
	GetUserData(data *StrategyData, code string)
}

type AuthGoCore struct {
	Data     StrategyData
	Strategy Strategy
}
