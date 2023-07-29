package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
)

type MunicipalityHandler struct {
	store stores.MunicipalityStore
}

func NewMunicipalityHandler(store stores.MunicipalityStore) *MunicipalityHandler {
	return &MunicipalityHandler{store: store}
}

func (handler MunicipalityHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	municipality, err := handler.store.FindByCode(vars["municipalityCode"])
	if err != nil {
		JSONResponse(w, ErrMunicipalityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, municipality, http.StatusOK)
}

func (handler MunicipalityHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}

func (handler MunicipalityHandler) ListByProvinceCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.ListByProvinceCode(vars["provinceCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}

func (handler MunicipalityHandler) ListByDistrictCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.ListByDistrictCode(vars["districtCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}
