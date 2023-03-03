package utils

import (
	"github.com/gin-gonic/gin"
)

func GetRequestProtocol(c *gin.Context) string {
		if c.Request.TLS != nil {
			return "https"
		} else {
			return "http"
		}
}