package models

import "encoding/json"

type Province struct {
	Id          int     `json:"-"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	IncomeClass string  `json:"income_class"`
	Population  int     `json:"population"`
	RegionCode  *string `json:"region_code"`
}

func (p Province) ID() int {
	return p.Id
}

func (p Province) String() string {
	data, _ := json.Marshal(p)

	return string(data)
}
