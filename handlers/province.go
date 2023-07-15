package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"ph-locations/stores"
	"strconv"
)

type ProvinceHandler struct {
	store stores.ProvinceStore
}

func NewProvinceHandler(store stores.ProvinceStore) *ProvinceHandler {
	return &ProvinceHandler{store: store}
}

func (handler ProvinceHandler) ShowProvinceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		JSONResponse(w, ResponseMessage{http.StatusBadRequest, "invalid province id"}, http.StatusBadRequest)
		return
	}

	province, err := handler.store.Find(id)

	if err != nil {
		JSONResponse(w, ResponseMessage{http.StatusNotFound, "province not found"}, http.StatusNotFound)
		return
	}

	JSONResponse(w, province, http.StatusOK)
}
