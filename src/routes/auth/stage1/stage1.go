package stage1

import (
	"fmt"
	"net/http"

	"auth-go/src/env"
	"auth-go/src/models/auth"
	"auth-go/src/models/auth/state"
	"auth-go/src/models/auth/templates"
	"auth-go/src/models/client/connection"
	"auth-go/src/utils"

	"github.com/gin-gonic/gin"
)



func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("app")
		provider := c.Param("provider")
		redirect := c.Query("redirect")

		clientMongo := connection.NewClientMongo(e.Mongo)
		client := clientMongo.GetByName(name)

		if (client == nil) {
			c.Data(http.StatusOK, "", []byte(`hello`))
			return
		}

		if client.CheckProvider(provider) == false {
			c.Data(http.StatusOK, "", []byte(`hello`))
			return
		}


	  providerTemplate := templates.ProviderTemplates[provider]

		stage1 := auth.Stage1 {
			Url: providerTemplate.Stage1.URL,
			ClientId: client.Providers[provider].ClientId,
			ResponseType: providerTemplate.Stage1.ResponseType,
			RedirectUri: fmt.Sprintf("%s://%s/auth/%s/stage2/%s", utils.GetRequestProtocol(c), c.Request.Host, name, provider),
			Scope: providerTemplate.Stage1.Scope,
			State: string(*state.StringifyState(&state.State {
				Redirect: redirect,
			})) ,
		}

		c.Redirect(302, stage1.GetUri())
	}
}
