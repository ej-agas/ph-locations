package handlers

import "net/http"

var (
	ErrInvalidPSGC             = ResponseMessage{StatusCode: http.StatusUnprocessableEntity, Message: "Invalid PSGC"}
	ErrInvalidLimit            = ResponseMessage{StatusCode: http.StatusUnprocessableEntity, Message: "Invalid Limit"}
	ErrInvalidPage             = ResponseMessage{StatusCode: http.StatusUnprocessableEntity, Message: "Invalid Page"}
	ErrRegionNotFound          = ResponseMessage{StatusCode: http.StatusNotFound, Message: "Region not found"}
	ErrProvinceNotFound        = ResponseMessage{StatusCode: http.StatusNotFound, Message: "Province not found"}
	ErrMunicipalityNotFound    = ResponseMessage{StatusCode: http.StatusNotFound, Message: "Municipality not found"}
	ErrDistrictNotFound        = ResponseMessage{StatusCode: http.StatusNotFound, Message: "District not found"}
	ErrCityNotFound            = ResponseMessage{StatusCode: http.StatusNotFound, Message: "City not found"}
	ErrBarangayNotFound        = ResponseMessage{StatusCode: http.StatusNotFound, Message: "Barangay not found"}
	ErrSubMunicipalityNotFound = ResponseMessage{StatusCode: http.StatusNotFound, Message: "Sub-Municipality not found"}
	ErrSGUNotFound             = ResponseMessage{StatusCode: http.StatusNotFound, Message: "Special Government Unit not found"}
)
