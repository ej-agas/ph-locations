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

// ShowByCode godoc
//
//	@summary		Show City
//	@description	Show City by Philippine Standard Geographic Code (PSGC)
//	@tags			cities
//	@produce		json
//	@success		200			{object}	models.City
//	@failure		404			{object}	handlers.ResponseMessage
//	@param			cityCode	path		string	true	"City's PSGC"
//	@router			/cities/{cityCode} [get]
func (handler CityHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city, err := handler.store.FindByCode(vars["cityCode"])
	if err != nil {
		JSONResponse(w, ErrCityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, city, http.StatusOK)
}

// List godoc
//
//	@summary		List Cities
//	@description	List Cities
//	@tags			cities
//	@produce		json
//	@success		200		{object}	stores.Collection[models.City]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@router			/cities [get]
func (handler CityHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	cities, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, cities, http.StatusOK)
}

// ListByProvinceCode godoc
//
//	@summary		List Cities
//	@description	List Cities by Province's Philippine Standard Geographic Code (PSGC)
//	@tags			cities
//	@produce		json
//	@success		200				{object}	stores.Collection[models.City]
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			provinceCode	path		string	true	"Province's PSGC"
//	@param			order			query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort			query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit			query		string	false	"Limit results per page. (default: 25)"
//	@param			page			query		string	false	"Page number. (default: 1)"
//	@router			/provinces/{provinceCode}/cities [get]
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

// ListByDistrictCode godoc
//
//	@summary		List Cities
//	@description	List Cities by District's Philippine Standard Geographic Code (PSGC)
//	@tags			cities
//	@produce		json
//	@success		200				{object}	stores.Collection[models.City]
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			districtCode	path		string	true	"District's PSGC"
//	@param			order			query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort			query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit			query		string	false	"Limit results per page. (default: 25)"
//	@param			page			query		string	false	"Page number. (default: 1)"
//	@router			/districts/{districtCode}/cities [get]
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
