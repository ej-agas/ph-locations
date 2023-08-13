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

// ShowByCode godoc
//
//	@summary		Show Special Government Unit
//	@description	Show Special Government Unit by Philippine Standard Geographic Code (PSGC)
//	@tags			special-government-units
//	@produce		json
//	@success		200		{object}	models.SpecialGovernmentUnit
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			sguCode	path		string	true	"Special Government Unit PSGC"
//	@router			/special-government-units/{sguCode} [get]
func (handler SGUHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sgu, err := handler.store.FindByCode(vars["sguCode"])
	if err != nil {
		JSONResponse(w, ErrSGUNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, sgu, http.StatusOK)
}

// List godoc
//
//	@summary		List Special Government Units
//	@description	List Special Government Units
//	@tags			special-government-units
//	@produce		json
//	@success		200		{object}	stores.Collection[models.SpecialGovernmentUnit]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@param			q		query		string	false	"Search by special government unit name"
//	@router			/special-government-units [get]
func (handler SGUHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	sgus, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, sgus, http.StatusOK)
}

// ListByProvinceCode godoc
//
//	@summary		List Special Government Units
//	@description	List Special Government Units by Province's Philippine Standard Geographic Code (PSGC)
//	@tags			special-government-units
//	@produce		json
//	@success		200				{object}	models.SpecialGovernmentUnit
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			provinceCode	path		string	true	"Province's PSGC"
//	@param			order			query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort			query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit			query		string	false	"Limit results per page. (default: 25)"
//	@param			page			query		string	false	"Page number. (default: 1)"
//	@param			q				query		string	false	"Search by special government unit name"
//	@router			/provinces/{provinceCode}/special-government-units [get]
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
