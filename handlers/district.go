package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
)

type DistrictHandler struct {
	store stores.DistrictStore
}

func NewDistrictHandler(store stores.DistrictStore) *DistrictHandler {
	return &DistrictHandler{store: store}
}

// ShowByCode godoc
//
//	@summary		Show District
//	@description	Show District by Philippine Standard Geographic Code (PSGC)
//	@tags			districts
//	@produce		json
//	@success		200				{object}	models.District
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			districtCode	path		string	true	"District's PSGC"
//	@router			/districts/{districtCode} [get]
func (handler DistrictHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	district, err := handler.store.FindByCode(vars["districtCode"])
	if err != nil {
		JSONResponse(w, ErrDistrictNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, district, http.StatusOK)
}

// List godoc
//
//	@summary		List Districts
//	@description	List Districts
//	@tags			districts
//	@produce		json
//	@success		200		{object}	stores.Collection[models.District]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@param			q		query		string	false	"Search by district name"
//	@router			/districts [get]
func (handler DistrictHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	districts, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, districts, http.StatusOK)
}

// ListByRegionCode godoc
//
//	@summary		List Districts By Region Code
//	@description	List Districts by Region's Philippine Standard Geographic Code (PSGC)
//	@tags			districts
//	@produce		json
//	@success		200			{object}	stores.Collection[models.District]
//	@failure		404			{object}	handlers.ResponseMessage
//	@param			regionCode	path		string	true	"Region's PSGC"
//	@param			order		query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort		query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit		query		string	false	"Limit results per page. (default: 25)"
//	@param			page		query		string	false	"Page number. (default: 1)"
//	@param			q			query		string	false	"Search by district name"
//	@router			/regions/{regionCode}/districts [get]
func (handler DistrictHandler) ListByRegionCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	districts, err := handler.store.ListByRegionCode(vars["regionCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, districts, http.StatusOK)
}
