package app

import (
	"github.com/deltrinos/fizz-buzz/engine"
	"github.com/deltrinos/fizz-buzz/interfaces"
	"github.com/deltrinos/fizz-buzz/store"
)

var (
	Engine *engine.Engine
	Store  *store.Store

	Queries interfaces.Fizzable
)

// this interface is used by endpoints to make queries from store
func SetQueries(q interfaces.Fizzable) {
	Queries = q
}
