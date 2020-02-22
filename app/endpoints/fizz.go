package endpoints

import (
	"github.com/deltrinos/fizz-buzz/app"
	"github.com/deltrinos/fizz-buzz/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func init() {
	app.Engine.Handler.GET("/fizz", func(c *gin.Context) {
		var params models.FizzParams
		err := c.ShouldBindQuery(&params)
		if err != nil {
			log.Debug().Msgf("failed to bind query %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		res := app.Queries.GetResults(params)
		if res != nil {
			c.JSON(http.StatusOK, res)
		} else {
			c.AbortWithStatus(http.StatusNoContent)
		}
	})
}
