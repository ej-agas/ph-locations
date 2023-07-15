package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "Philippine Standard Geographic Code (PSGC) REST API",
	}

	JSONResponse(w, body, http.StatusOK)
}
