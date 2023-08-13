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

// ShowByCode godoc
//
//	@summary		Show Sub-Municipality
//	@description	Show Sub-Municipality by Philippine Standard Geographic Code (PSGC)
//	@tags			sub-municipalities
//	@produce		json
//	@success		200					{object}	models.SubMunicipality
//	@failure		404					{object}	handlers.ResponseMessage
//	@param			subMunicipalityCode	path		string	true	"Sub-Municipality's PSGC"
//	@router			/sub-municipalities/{subMunicipalityCode} [get]
func (handler SubMunicipalityHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city, err := handler.store.FindByCode(vars["subMunicipalityCode"])
	if err != nil {
		JSONResponse(w, ErrSubMunicipalityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, city, http.StatusOK)
}

// List godoc
//
//	@summary		List Sub-Municipalities
//	@description	List Sub-Municipalities
//	@tags			sub-municipalities
//	@produce		json
//	@success		200		{object}	stores.Collection[models.SubMunicipality]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@param			q		query		string	false	"Search by sub-municipality name"
//	@router			/sub-municipalities [get]
func (handler SubMunicipalityHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	subMunicipalities, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, subMunicipalities, http.StatusOK)
}

// ListByCityCode godoc
//
//	@summary		List Sub-Municipalities
//	@description	List Sub-Municipalities by City's Philippine Standard Geographic Code (PSGC)
//	@tags			sub-municipalities
//	@produce		json
//	@success		200			{object}	stores.Collection[models.SubMunicipality]
//	@failure		404			{object}	handlers.ResponseMessage
//	@param			cityCode	path		string	true	"City's PSGC"
//	@param			order		query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort		query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit		query		string	false	"Limit results per page. (default: 25)"
//	@param			page		query		string	false	"Page number. (default: 1)"
//	@param			q			query		string	false	"Search by sub-municipality name"
//	@router			/cities/{cityCode}/sub-municipalities [get]
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
