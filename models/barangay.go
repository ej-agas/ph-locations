package models

type Barangay struct {
	Id                int    `json:"-"`
	Code              string `json:"code"`
	Name              string `json:"name"`
	UrbanRural        string `json:"urban_rural"`
	Population        int    `json:"population"`
	CityId            int    `json:"city_id,omitempty"`
	MunicipalityId    int    `json:"municipality_id,omitempty"`
	SubMunicipalityId int    `json:"sub_municipality_id,omitempty"`
}
