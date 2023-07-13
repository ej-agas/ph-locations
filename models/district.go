package models

type District struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Population int    `json:"population"`
	RegionId   int    `json:"region_id"`
}
