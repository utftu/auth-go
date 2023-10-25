package user

import (
	"encoding/json"
	"net/http"
	"service/src/env"

	userConnection "service/src/models/auth/user/connection"

	"github.com/gin-gonic/gin"
)

func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		userCode := c.Query("code")

		userMongo := userConnection.NewUserMongoConnection(e.Mongo)
		user := userMongo.Get("code", userCode)

		if user == nil {
			c.Data(http.StatusOK, "", []byte(`invalid code`))
			return
		}

		json, _ := json.Marshal(&user.User)

		c.Data(200, "application/json", json)
	}
}
