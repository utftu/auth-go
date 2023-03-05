package stage2

import (
	"auth-go/src/env"
	"auth-go/src/models/auth"
	"auth-go/src/models/auth/state"
	"auth-go/src/models/auth/templates"
	userConnection "auth-go/src/models/auth/user/connection"
	"auth-go/src/models/client/connection"
	"auth-go/src/utils"

	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("app")
		provider := c.Param("provider")
		code := c.Query("code")
		stateQuery := c.Query("state")

		fmt.Println("-----", "code", code)

		parsedState := state.ParseState(stateQuery)
		if parsedState == nil {
			c.Data(http.StatusOK, "", []byte(`Invalid state`))
			return
		}

		if code == "" {
			c.Data(http.StatusOK, "", []byte(`Invalid code`))
			return
		}

		clientMongo := connection.NewClientMongo(e.Mongo)
		client := clientMongo.GetByName(name)

		if client == nil {
			c.Data(http.StatusOK, "", []byte(`Invalid client`))
			return
		}

		if client.CheckProvider(provider) == false {
			c.Data(http.StatusOK, "", []byte(`Invalid provider`))
			return
		}

		providerTemplate := templates.ProviderTemplates[provider]

		stage2 := auth.Stage2{
			Url:  providerTemplate.Stage2.URL,
			Name: provider,

			ClientId:     client.Providers[provider].ClientId,
			ClientSecret: client.Providers[provider].ClientSecret,
			GrantType:    providerTemplate.Stage2.GrantType,
			RedirectUri:  fmt.Sprintf("%s://%s/auth/%s/stage2/%s", utils.GetRequestProtocol(c), c.Request.Host, name, provider),
			Code:         code,
			State:        stateQuery,
		}

		user, error := stage2.Request()

		if error != nil {
			c.Data(http.StatusOK, "", []byte(`Request to provider failed`))
			return
		}

		userMongoConnection := userConnection.NewUserMongoConnection(e.Mongo)
		userCode := userMongoConnection.Save(user)

		parsedRedirect, err := url.Parse(parsedState.Redirect)
		if err != nil {
			c.Data(http.StatusOK, "", []byte(`Invalid redirect`))
			return
		}

		parsedRedirectQuery := parsedRedirect.Query()
		parsedRedirectQuery.Set("code", userCode)
		parsedRedirect.RawQuery = parsedRedirectQuery.Encode()

		c.Redirect(302, parsedRedirect.String())
	}
}
