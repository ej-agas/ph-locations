package models

type Province struct {
	Id          int    `json:"-"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	IncomeClass string `json:"income_class"`
	Population  int    `json:"population"`
	RegionId    int    `json:"region_id"`
}
