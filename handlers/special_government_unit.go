package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
)

type SGUHandler struct {
	store stores.SpecialGovernmentUnit
}

func NewSGUHandler(store stores.SpecialGovernmentUnit) *SGUHandler {
	return &SGUHandler{store: store}
}

func (handler SGUHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sgu, err := handler.store.FindByCode(vars["sguCode"])
	if err != nil {
		JSONResponse(w, ErrSGUNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, sgu, http.StatusOK)
}

func (handler SGUHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	sgus, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, sgus, http.StatusOK)
}

func (handler SGUHandler) ListByProvinceCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	sgus, err := handler.store.ListByProvinceCode(vars["provinceCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, sgus, http.StatusOK)
}
