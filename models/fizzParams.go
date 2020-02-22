package models

type FizzParams struct {
	Int1  int64  `form:"int1" binding:"required"`
	Int2  int64  `form:"int2" binding:"required"`
	Limit int64  `form:"limit" binding:"required"`
	Str1  string `form:"str1" binding:"required"`
	Str2  string `form:"str2" binding:"required"`
}
