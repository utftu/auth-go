package auth

type Data struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

type Strategy interface {
	GetUserRedirect(data *Data) string
	GetUserData(data *Data, code string)
}

type AuthGoCore struct {
	Data     Data
	Strategy Strategy
}
