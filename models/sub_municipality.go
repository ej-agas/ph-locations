package models

import "encoding/json"

type SubMunicipality struct {
	Id         int     `json:"-"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Population int     `json:"population"`
	CityCode   *string `json:"city_code"`
}

func (s SubMunicipality) ID() int {
	return s.Id
}

func (s SubMunicipality) String() string {
	data, _ := json.Marshal(s)

	return string(data)
}
