package models

type SpecialGovernmentUnit struct {
	Id         int    `json:"-"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	ProvinceId *int   `json:"province_id"`
}
