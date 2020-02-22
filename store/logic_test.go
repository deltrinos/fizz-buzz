package store

import (
	"github.com/deltrinos/fizz-buzz/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIntegerStuff(t *testing.T) {
	Convey("FizzParamsToStrArray", t, func() {
		tests := []struct {
			p      models.FizzParams
			except []string
		}{
			{
				p: models.FizzParams{
					Int1:  2,
					Int2:  3,
					Limit: 1,
					Str1:  "toto",
					Str2:  "titi",
				},
				except: []string{"1"},
			},
			{
				p: models.FizzParams{
					Int1:  2,
					Int2:  3,
					Limit: -3,
					Str1:  "toto",
					Str2:  "titi",
				},
				except: []string{"1", "tototiti", "-1", "toto", "titi"},
			},
			{
				p: models.FizzParams{
					Int1:  5,
					Int2:  6,
					Limit: 10,
					Str1:  "toto",
					Str2:  "titi",
				},
				except: []string{"1", "2", "3", "4", "toto", "titi", "7", "8", "9", "toto"},
			},
		}

		for _, test := range tests {
			res := FizzParamsToStrArray(test.p)
			So(res, ShouldResemble, test.except)
		}
	})
}
