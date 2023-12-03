package start

import (
	"fmt"
	"net/http"
	"os"

	"core"
	"site/src/env"
	"site/src/models/auth"
	"site/src/models/client"

	"github.com/gin-gonic/gin"
)

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("client")
		provider := c.Param("provider")
		redirect := c.Query("redirect")

		clientMongo := client.NewClientMongo(e.Mongo)
		client := clientMongo.GetByName(name)

		if client == nil {
			c.Data(http.StatusOK, "", []byte(`invalid app name`))
			return
		}

		if client.CheckProvider(provider) == false {
			c.Data(http.StatusOK, "", []byte(`invalid provider`))
			return
		}

		// providerTemplate := templates.ProviderTemplates[provider]

		strategy := auth.SelectStrategy(provider, &core.StrategyData{
			ClientId: client.Providers[provider].ClientId,
			ClientSecret: client.Providers[provider].ClientSecret,
			RedirectUrl: redirect,
			ServiceRedirectUrl: fmt.Sprintf("%s/auth/%s/stage2/%s", os.Getenv("EXTERNAL_URL"), name, provider),
		})

		c.Redirect(302, strategy.GetUserRedirectUrl())
	}
}
