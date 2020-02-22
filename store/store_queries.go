package store

import (
	"github.com/deltrinos/fizz-buzz/models"
	"github.com/rs/zerolog/log"
)

func (s *Store) GetResults(params models.FizzParams) []string {

	// serve from cache
	qRes, qExists := queriesCache.Get(params)
	if qExists {
		log.Debug().Msgf("response from cache")
		return qRes
	}

	res := FizzParamsToStrArray(params)

	// save to cache and returns
	return queriesCache.Save(params, res)
}

func (s *Store) Stats() interface{} {
	return s.RequestURIs
}
