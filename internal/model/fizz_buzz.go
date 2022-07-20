package model

type Parameters struct {
	FirstInteger  int `json:"int1"`
	SecondInteger int `json:"int2"`
	Limit         int `json:"limit"`

	FirstString  string `json:"str1"`
	SecondString string `json:"str2"`
}
