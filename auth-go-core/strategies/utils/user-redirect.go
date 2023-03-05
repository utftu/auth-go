package utils

import (
	"encoding/json"
	"net/url"
	"strings"
)

type state struct {
	Redirect string `json:"redirect"`
}

type UserRedirectData struct {
	Url string

	ClientId         string
	ResponseType     string
	RedirectUrl      string
	UserRedirectUrl string
	Scope            []string
}

func GetUserRedirectUrl(data *UserRedirectData) string {
	parsedUrl, err := url.Parse(data.Url)
	if err != nil {
		return ""
	}
	params := parsedUrl.Query()

	stateInstance := state{
		Redirect: data.RedirectUrl,
	}
	stateBytes, _ := json.Marshal(&stateInstance)
	params.Add("client_id", data.ClientId)
	params.Add("response_type", data.ResponseType)
	params.Add("redirect_uri", data.UserRedirectUrl)
	params.Add("scope", strings.Join(data.Scope, " "))
	params.Add("state", string(stateBytes))
	parsedUrl.RawQuery = params.Encode()
	return parsedUrl.String()
}
