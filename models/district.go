package models

import "encoding/json"

type District struct {
	Id         int     `json:"-"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Population int     `json:"population"`
	RegionCode *string `json:"region_code"`
}

func (d District) ID() int {
	return d.Id
}

func (d District) String() string {
	data, _ := json.Marshal(d)

	return string(data)
}
