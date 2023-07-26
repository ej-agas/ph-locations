package models

type SpecialGovernmentUnit struct {
	Id           int     `json:"-"`
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	ProvinceCode *string `json:"province_codes"`
}
