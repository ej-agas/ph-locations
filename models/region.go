package models

type Region struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}
