package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
)

type SubMunicipalityHandler struct {
	store stores.SubMunicipality
}

func NewSubMunicipalityHandler(store stores.SubMunicipality) *SubMunicipalityHandler {
	return &SubMunicipalityHandler{store: store}
}

func (handler SubMunicipalityHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city, err := handler.store.FindByCode(vars["subMunicipalityCode"])
	if err != nil {
		JSONResponse(w, ErrSubMunicipalityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, city, http.StatusOK)
}

func (handler SubMunicipalityHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	subMunicipalities, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, subMunicipalities, http.StatusOK)
}

func (handler SubMunicipalityHandler) ListByCityCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	subMunicipalities, err := handler.store.ListByCityCode(vars["cityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, subMunicipalities, http.StatusOK)
}
