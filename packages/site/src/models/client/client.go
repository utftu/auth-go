package client

import (
	"regexp"
	"fmt"
)

type ClientProvider struct {
	Name         string
	ClientId     string `bson:"client_id"`
	ClientSecret string `bson:"client_secret"`
}

type Client struct {
	Id          string `bson:"_id"`
	Name        string
	AllowedUrls []string `bson:"allowed_urls"`

	Providers map[string]ClientProvider
}

func (client *Client) CheckProvider(providerName string) bool {
	for _, provider := range client.Providers {
		if provider.Name == providerName {
			return true
		}
	}
	return false
}

func (client *Client) CheckAllowedUrl(url string) bool {
	fmt.Println("-----", "check", client);
	for _, allowedUrl := range client.AllowedUrls {
		fmt.Println("-----", "allowedUrl", allowedUrl);
		matched, _ := regexp.MatchString(allowedUrl, url)
		if matched {
			return true
		}
	}
	return false
}
