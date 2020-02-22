package interfaces

import (
	"github.com/deltrinos/fizz-buzz/models"
)

type Fizzable interface {
	GetResults(params models.FizzParams) []string
	Stats() interface{}
}
