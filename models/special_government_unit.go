package models

import "encoding/json"

type SpecialGovernmentUnit struct {
	Id           int     `json:"-"`
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	ProvinceCode *string `json:"province_codes"`
}

func (s SpecialGovernmentUnit) ID() int {
	return s.Id
}

func (s SpecialGovernmentUnit) String() string {
	data, _ := json.Marshal(s)

	return string(data)
}
