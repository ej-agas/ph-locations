package models

type Municipality struct {
	Id           int     `json:"-"`
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	IncomeClass  string  `json:"income_class"`
	Population   int     `json:"population"`
	ProvinceCode *string `json:"province_code,omitempty"`
	DistrictCode *string `json:"district_code,omitempty"`
}
