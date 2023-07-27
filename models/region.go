package models

import "encoding/json"

type Region struct {
	Id         int    `json:"-"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func (r Region) ID() int {
	return r.Id
}

func (r Region) String() string {
	data, _ := json.Marshal(r)

	return string(data)
}
