package engine

import (
	"github.com/deltrinos/fizz-buzz/app"
	"github.com/deltrinos/fizz-buzz/app/conf"
	"github.com/deltrinos/fizz-buzz/engine"
)

func init() {
	app.Engine = engine.NewEngine(conf.Env.Addr)
}
