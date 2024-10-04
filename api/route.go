package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/simple-api-xmu/config"
)

func Run() {

	config.Init()
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := app.Group("/api/v1")

	api.Use(BasicAuth())

	NewHandler(api, NewStore(config.NewDB(config.ENV.DB)))

	app.Run(config.ENV.Address())
}
