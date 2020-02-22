package middlewares

import "github.com/deltrinos/fizz-buzz/app"

func init() {
	app.Engine.Handler.Use(RequestURISaver)
}
