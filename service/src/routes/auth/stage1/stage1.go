package stage1

import (
	"fmt"
	"net/http"
	"os"

	"service/src/env"
	"service/src/models/auth"
	"service/src/models/client/connection"

	"auth-go-core"

	"github.com/gin-gonic/gin"
)

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("app")
		provider := c.Param("provider")
		redirect := c.Query("redirect")

		clientMongo := connection.NewClientMongo(e.Mongo)
		client := clientMongo.GetByName(name)

		if client == nil {
			c.Data(http.StatusOK, "", []byte(`hello`))
			return
		}

		if client.CheckProvider(provider) == false {
			c.Data(http.StatusOK, "", []byte(`hello`))
			return
		}

		// providerTemplate := templates.ProviderTemplates[provider]

		strategy := auth.SelectStrategy(provider, &authGoCore.StrategyData{
			ClientId: client.Providers[provider].ClientId,
			ClientSecret: client.Providers[provider].ClientSecret,
			RedirectUrl: redirect,
			ServiceRedirectUrl: fmt.Sprintf("%s/auth/%s/stage2/%s", os.Getenv("EXTERNAL_URL"), name, provider),
		})

		c.Redirect(302, strategy.GetUserRedirectUrl())

		// stage1 := auth.Stage1{
		// 	Url:          providerTemplate.Stage1.URL,
		// 	ClientId:     client.Providers[provider].ClientId,
		// 	ResponseType: providerTemplate.Stage1.ResponseType,
		// 	RedirectUri:  fmt.Sprintf("%s://%s/auth/%s/stage2/%s", utils.GetRequestProtocol(c), c.Request.Host, name, provider),
		// 	Scope:        providerTemplate.Stage1.Scope,
		// 	State: string(*state.StringifyState(&state.State{
		// 		Redirect: redirect,
		// 	})),
		// }

		// c.Redirect(302, stage1.GetUri())
	}
}
