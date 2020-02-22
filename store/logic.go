package store

import (
	"fmt"
	"github.com/deltrinos/fizz-buzz/models"
)

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func FizzParamsToStrArray(params models.FizzParams) []string {
	len := Abs(params.Limit-1) + 1
	res := make([]string, len)
	diff := int64(1)
	if params.Limit < 1 {
		diff = int64(-1)
	}

	idx := 0
	str12 := params.Str1 + params.Str2
	limit := params.Limit + diff
	for start := int64(1); start != limit; start += diff {
		b1 := start%params.Int1 == 0
		b2 := start%params.Int2 == 0

		if b1 && b2 {
			res[idx] = str12
		} else if b1 {
			res[idx] = params.Str1
		} else if b2 {
			res[idx] = params.Str2
		} else {
			res[idx] = fmt.Sprintf("%d", start)
		}
		idx++
	}
	return res
}
