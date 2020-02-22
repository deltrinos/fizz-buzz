package middlewares

import (
	"github.com/deltrinos/fizz-buzz/app"
	"github.com/gin-gonic/gin"
)

func RequestURISaver(c *gin.Context) {
	app.Store.Save(c.Request.RequestURI)
	c.Next()
}
