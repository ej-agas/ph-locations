package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
)

type CityHandler struct {
	store stores.CityStore
}

func NewCityHandler(store stores.CityStore) *CityHandler {
	return &CityHandler{store: store}
}

func (handler CityHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city, err := handler.store.FindByCode(vars["cityCode"])
	if err != nil {
		JSONResponse(w, ErrCityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, city, http.StatusOK)
}

func (handler CityHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	cities, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, cities, http.StatusOK)
}

func (handler CityHandler) ListByProvinceCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	districts, err := handler.store.ListByProvinceCode(vars["provinceCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, districts, http.StatusOK)
}

func (handler CityHandler) ListByDistrictCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	districts, err := handler.store.ListByDistrictCode(vars["districtCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, districts, http.StatusOK)
}
