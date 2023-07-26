package models

type District struct {
	Id         int     `json:"-"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Population int     `json:"population"`
	RegionCode *string `json:"region_code"`
}
