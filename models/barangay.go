package models

type Barangay struct {
	Id                        int     `json:"-"`
	Code                      string  `json:"code"`
	Name                      string  `json:"name"`
	UrbanRural                string  `json:"urban_rural"`
	Population                int     `json:"population"`
	CityCode                  *string `json:"city_code,omitempty"`
	MunicipalityCode          *string `json:"municipality_code,omitempty"`
	SubMunicipalityCode       *string `json:"sub_municipality_code,omitempty"`
	SpecialGovernmentUnitCode *string `json:"special_government_unit_code,omitempty"`
}
