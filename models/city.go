package models

type City struct {
	Id          int    `json:"-"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	CityClass   string `json:"city_class"`
	IncomeClass string `json:"income_class"`
	Population  int    `json:"population"`
	ProvinceId  *int   `json:"province_id,omitempty"`
	DistrictId  *int   `json:"district_id,omitempty"`
}
