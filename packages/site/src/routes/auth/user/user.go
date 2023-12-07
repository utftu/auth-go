package user

import (
	"encoding/json"
	"net/http"
	"site/src/env"

	"site/src/models/auth/user"

	"github.com/gin-gonic/gin"
)

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		userCode := c.Query("code")

		userMongo := user.NewUserMongoConnection(e.Mongo)
		user := userMongo.Get("code", userCode)

		if user == nil {
			c.Data(http.StatusOK, "", []byte(`invalid code`))
			return
		}

		json, _ := json.Marshal(&user.User)

		c.Data(200, "application/json", json)
	}
}
