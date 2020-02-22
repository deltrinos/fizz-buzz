package endpoints

import (
	"github.com/deltrinos/fizz-buzz/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	app.Engine.Handler.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, app.Queries.Stats())
	})
}
