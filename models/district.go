package models

type District struct {
	Id         int    `json:"-"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Population int    `json:"population"`
	RegionId   *int   `json:"region_id"`
}
