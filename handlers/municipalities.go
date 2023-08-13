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

// ShowByCode godoc
//
//	@summary		Show Municipality
//	@description	Show Municipality by Philippine Standard Geographic Code (PSGC)
//	@tags			municipalities
//	@produce		json
//	@success		200					{object}	models.Municipality
//	@failure		404					{object}	handlers.ResponseMessage
//	@param			municipalityCode	path		string	true	"Municipality's PSGC"
//	@router			/municipalities/{municipalityCode} [get]
func (handler MunicipalityHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	municipality, err := handler.store.FindByCode(vars["municipalityCode"])
	if err != nil {
		JSONResponse(w, ErrMunicipalityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, municipality, http.StatusOK)
}

// List godoc
//
//	@summary		List Municipalities
//	@description	List Municipalities
//	@tags			municipalities
//	@produce		json
//	@success		200		{object}	stores.Collection[models.Municipality]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@param			q		query		string	false	"Search by municipality name"
//	@router			/municipalities [get]
func (handler MunicipalityHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}

// ListByProvinceCode godoc
//
//	@summary		List Municipalities
//	@description	List Municipalities by Province's Philippine Standard Geographic Code (PSGC)
//	@tags			municipalities
//	@produce		json
//	@success		200				{object}	stores.Collection[models.Municipality]
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			provinceCode	path		string	true	"Province's PSGC"
//	@param			order			query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort			query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit			query		string	false	"Limit results per page. (default: 25)"
//	@param			page			query		string	false	"Page number. (default: 1)"
//	@router			/provinces/{provinceCode}/municipalities [get]
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

// ListByDistrictCode godoc
//
//	@summary		List Municipalities
//	@description	List Municipalities by District's Philippine Standard Geographic Code (PSGC)
//	@tags			municipalities
//	@produce		json
//	@success		200				{object}	stores.Collection[models.Municipality]
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			districtCode	path		string	true	"District's PSGC"
//	@param			order			query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort			query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit			query		string	false	"Limit results per page. (default: 25)"
//	@param			page			query		string	false	"Page number. (default: 1)"
//	@router			/districts/{districtCode}/municipalities [get]
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
