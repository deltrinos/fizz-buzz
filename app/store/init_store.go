package store

import (
	"github.com/deltrinos/fizz-buzz/app"
	"github.com/deltrinos/fizz-buzz/store"
)

func init() {
	app.Store = store.NewStore().Run()

	// make store serve queries
	app.SetQueries(app.Store)
}
