package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
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
		JSONResponse(w, ErrInvalidPSGC, http.StatusUnprocessableEntity)
		return
	}

	province, err := handler.store.Find(id)

	if err != nil {
		JSONResponse(w, ErrProvinceNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, province, http.StatusOK)
}

func (handler ProvinceHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	province, err := handler.store.FindByCode(r.Context(), vars["provinceCode"])
	if err != nil {
		JSONResponse(w, ErrProvinceNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, province, http.StatusOK)
}

func (handler ProvinceHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.List(*opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}

func (handler ProvinceHandler) ListByRegionId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.ListByRegionCode(vars["regionCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}
