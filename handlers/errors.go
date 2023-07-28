package handlers

import "net/http"

var (
	ErrInvalidPSGC      = ResponseMessage{StatusCode: http.StatusUnprocessableEntity, Message: "Invalid PSGC"}
	ErrInvalidLimit     = ResponseMessage{StatusCode: http.StatusUnprocessableEntity, Message: "Invalid Limit"}
	ErrInvalidPage      = ResponseMessage{StatusCode: http.StatusUnprocessableEntity, Message: "Invalid Page"}
	ErrProvinceNotFound = ResponseMessage{StatusCode: http.StatusNotFound, Message: "Province not found"}
	ErrDistrictNotFound = ResponseMessage{StatusCode: http.StatusNotFound, Message: "District not found"}
)
