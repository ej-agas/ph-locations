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

// ShowByCode godoc
//
//	@summary		Show Province
//	@description	Show Province by Philippine Standard Geographic Code (PSGC)
//	@tags			provinces
//	@produce		json
//	@success		200				{object}	models.Province
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			provinceCode	path		string	true	"Province's PSGC"
//	@router			/provinces/{provinceCode} [get]
func (handler ProvinceHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	province, err := handler.store.FindByCode(r.Context(), vars["provinceCode"])
	if err != nil {
		JSONResponse(w, ErrProvinceNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, province, http.StatusOK)
}

// List godoc
//
//	@summary		List Provinces
//	@description	List Provinces
//	@tags			provinces
//	@produce		json
//	@success		200		{object}	stores.Collection[models.Province]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@router			/provinces [get]
func (handler ProvinceHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.List(*opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}

// ListByRegionCode godoc
//
//	@summary		List Provinces by Region Code
//	@description	List Provinces by Region's Philippine Standard Geographic Code (PSGC)
//	@tags			provinces
//	@produce		json
//	@success		200			{object}	stores.Collection[models.Province]
//	@failure		404			{object}	handlers.ResponseMessage
//	@param			regionCode	path		string	true	"Region's PSGC"
//	@param			order		query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort		query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit		query		string	false	"Limit results per page. (default: 25)"
//	@param			page		query		string	false	"Page number. (default: 1)"
//	@router			/regions/{regionCode}/provinces [get]
func (handler ProvinceHandler) ListByRegionCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	provinces, err := handler.store.ListByRegionCode(vars["regionCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}
