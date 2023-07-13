package models

type SubMunicipality struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Population int    `json:"population"`
	CityId     int    `json:"city_id"`
}
