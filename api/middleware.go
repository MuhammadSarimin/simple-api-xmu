package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/simple-api-xmu/types"
	log "github.com/sirupsen/logrus"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// check basic auth
		user, pass, ok := c.Request.BasicAuth()
		if !ok || user != "simple-api" || pass != "xmu" {
			log.Error("Unauthorized", fmt.Errorf("unauthorized"))
			c.JSON(401, types.Response{
				ResponseCode:    "401",
				ResponseMessage: "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
