package stage2

import (
	"os"
	"service/src/env"
	"service/src/models/auth"
	"service/src/models/auth/state"

	"service/src/models/auth/user/connection"
	"service/src/models/client/connection"

	"auth-go-core"

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

		strategy := auth.SelectStrategy(provider, &authGoCore.StrategyData{
			ClientId: client.Providers[provider].ClientId,
			ClientSecret: client.Providers[provider].ClientSecret,
			RedirectUrl: parsedState.Redirect,
			ServiceRedirectUrl: fmt.Sprintf("%s/auth/%s/stage2/%s", os.Getenv("EXTERNAL_URL"), name, provider),
		})

		user, err := strategy.GetUserData(code)

		if err != nil {
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
