package user

import (
	"auth-go/src/env"
	"encoding/json"
	"fmt"

	userConnection "auth-go/src/models/auth/user/connection"

	"github.com/gin-gonic/gin"
)


func CreateHandler(e *env.Env) func(c *gin.Context) {
	return func(c *gin.Context) {
		userCode := c.Query("code")
		
		userMongo := userConnection.NewUserMongoConnection(e.Mongo)
		user := userMongo.Get("code", userCode)

		fmt.Println("-----", "user.User", user);

		json, err := json.Marshal(&user)

		fmt.Println("-----", "err", err);

		c.Data(200, "application/json", json)

		// c.JSON(200, user)
	}
}