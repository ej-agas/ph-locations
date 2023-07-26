package models

type SubMunicipality struct {
	Id         int     `json:"-"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Population int     `json:"population"`
	CityCode   *string `json:"city_code"`
}
