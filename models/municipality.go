package models

type Municipality struct {
	Id          int    `json:"-"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	IncomeClass string `json:"income_class"`
	Population  int    `json:"population"`
	ProvinceId  int    `json:"province_id,omitempty"`
	DistrictId  int    `json:"district_id,omitempty"`
}
