package auth

import (
	// "net/http"
	"net/url"
	"strings"
)

type Stage1 struct {
	Url string

	ClientId     string
	ResponseType string
	RedirectUri  string
	Scope        []string
	State        string
}

func (enter *Stage1) GetUri() string {
	parsedUrl, err := url.Parse(enter.Url)
	if err != nil {
		return ""
	}
	params := parsedUrl.Query()
	params.Add("client_id", enter.ClientId)
	params.Add("response_type", enter.ResponseType)
	params.Add("redirect_uri", enter.RedirectUri)
	params.Add("scope", strings.Join(enter.Scope, " "))
	params.Add("state", enter.State)
	parsedUrl.RawQuery = params.Encode()
	return parsedUrl.String()
}
