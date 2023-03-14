package stage0

import (
	"fmt"
	"net/http"

	"service/src/env"
	"service/src/models/client"
	"service/src/models/client/connection"

	"github.com/gin-gonic/gin"
)

type htmlProps struct {
	Client      *client.Client
	RedirectUrl string
}

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("app")

		redirectUrl := c.Query("redirect")
		clientMongo := connection.NewClientMongo(e.Mongo)

		client := clientMongo.GetByName(name)

		if client == nil {
			c.Data(http.StatusOK, "", []byte(`hello`))
			return
		}

		if client.CheckAllowedUrl(redirectUrl) == false {
			c.Data(http.StatusOK, "", []byte(`hello123`))
			return
		}

		c.HTML(http.StatusOK, "stage0.html", &htmlProps{
			RedirectUrl: redirectUrl,
			Client:      client,
		})
	}
}
