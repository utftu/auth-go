package stage0

import (
	"fmt"
	"net/http"

	"auth-go/src/env"
	"auth-go/src/models/client"
	"auth-go/src/models/client/connection"

	"github.com/gin-gonic/gin"
)

type htmlProps struct {
	Client *client.Client
	RedirectUrl string
}

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("app")

		redirectUrl := c.Query("redirect")
		clientMongo := connection.NewClientMongo(e.Mongo)

		fmt.Println("name", name)


		client := clientMongo.GetByName(name)
		fmt.Println("client", client)

		if (client == nil) {
			c.Data(http.StatusOK, "", []byte(`hello`))
			return
		}

		if (client.CheckAllowedUrl(redirectUrl) == false) {
			c.Data(http.StatusOK, "", []byte(`hello123`))
			return
		}

		c.HTML(http.StatusOK, "stage0.html", &htmlProps {
      RedirectUrl: redirectUrl,
			Client: client,
		})
	}
}
