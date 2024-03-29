package client

import (
	"net/http"

	"site/src/env"
	"site/src/models/client"

	"github.com/gin-gonic/gin"
)

type htmlProps struct {
	Client      *client.Client
	RedirectUrl string
}

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("client")

		redirectUrl := c.Query("redirect")
		clientMongo := client.NewClientMongo(e.Mongo)

		client := clientMongo.GetByName(name)

		if client == nil {
			c.Data(http.StatusOK, "", []byte(`hello`))
			return
		}

		if client.CheckAllowedUrl(redirectUrl) == false {
			c.Data(http.StatusOK, "", []byte(`redirect url not allowed`))
			return
		}

		c.HTML(http.StatusOK, "client.html", &htmlProps{
			RedirectUrl: redirectUrl,
			Client:      client,
		})
	}
}
