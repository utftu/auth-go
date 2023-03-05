package templates

import (
	"encoding/json"
	"os"
)

func init() {
	jsonTemplate, _ := os.ReadFile("./src/models/auth/templates/provider-templates.json")
	json.Unmarshal(jsonTemplate, &ProviderTemplates)
}

type ProviderTemplate struct {
	Name   string `json:"name"`
	Stage1 struct {
		URL          string   `json:"url"`
		ResponseType string   `json:"responseType"`
		Scope        []string `json:"scope"`
	} `json:"stage1"`
	Stage2 struct {
		URL       string `json:"url"`
		GrantType string `json:"grantType"`
	} `json:"stage2"`
}

var ProviderTemplates map[string]ProviderTemplate
