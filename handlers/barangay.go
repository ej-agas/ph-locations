package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
)

type BarangayHandler struct {
	store stores.BarangayStore
}

func NewBarangayHandler(store stores.BarangayStore) *BarangayHandler {
	return &BarangayHandler{store: store}
}

// ShowByCode godoc
//
//	@summary		Show Barangay
//	@description	Show Barangay by Philippine Standard Geographic Code (PSGC)
//	@tags			barangays
//	@produce		json
//	@success		200				{object}	models.Barangay
//	@failure		404				{object}	handlers.ResponseMessage
//	@param			barangayCode	path		string	true	"Barangay's PSGC"
//	@router			/barangays/{barangayCode} [get]
func (handler BarangayHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	barangay, err := handler.store.FindByCode(vars["barangayCode"])
	if err != nil {
		JSONResponse(w, ErrBarangayNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, barangay, http.StatusOK)
}

// List godoc
//
//	@summary		List Barangays
//	@description	List Barangays
//	@tags			barangays
//	@produce		json
//	@success		200		{object}	stores.Collection[models.Barangay]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@param			q		query		string	false	"Search by barangay name"
//	@router			/barangays [get]
func (handler BarangayHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	barangays, err := handler.store.List(*opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}

// ListByCityCode godoc
//
//	@summary		List Barangays
//	@description	List Barangays by City's Philippine Standard Geographic Code (PSGC)
//	@tags			barangays
//	@produce		json
//	@success		200			{object}	stores.Collection[models.Barangay]
//	@failure		404			{object}	handlers.ResponseMessage
//	@param			cityCode	path		string	true	"City's PSGC"
//	@param			order		query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort		query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit		query		string	false	"Limit results per page. (default: 25)"
//	@param			page		query		string	false	"Page number. (default: 1)"
//	@param			q			query		string	false	"Search by barangay name"
//	@router			/cities/{cityCode}/barangays [get]
func (handler BarangayHandler) ListByCityCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	barangays, err := handler.store.ListByCityCode(vars["cityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}

// ListByMunicipalityCode godoc
//
//	@summary		List Barangays
//	@description	List Barangays by Municipality's Philippine Standard Geographic Code (PSGC)
//	@tags			barangays
//	@produce		json
//	@success		200					{object}	stores.Collection[models.Barangay]
//	@failure		404					{object}	handlers.ResponseMessage
//	@param			municipalityCode	path		string	true	"City's PSGC"
//	@param			order				query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort				query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit				query		string	false	"Limit results per page. (default: 25)"
//	@param			page				query		string	false	"Page number. (default: 1)"
//	@param			q					query		string	false	"Search by barangay name"
//	@router			/municipalities/{municipalityCode}/barangays [get]
func (handler BarangayHandler) ListByMunicipalityCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	barangays, err := handler.store.ListByMunicipalityCode(vars["municipalityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}

// ListBySubMunicipalityCode godoc
//
//	@summary		List Barangays
//	@description	List Barangays by Sub-Municipality's Philippine Standard Geographic Code (PSGC)
//	@tags			barangays
//	@produce		json
//	@success		200					{object}	stores.Collection[models.Barangay]
//	@failure		404					{object}	handlers.ResponseMessage
//	@param			subMunicipalityCode	path		string	true	"City's PSGC"
//	@param			order				query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort				query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit				query		string	false	"Limit results per page. (default: 25)"
//	@param			page				query		string	false	"Page number. (default: 1)"
//	@param			q					query		string	false	"Search by barangay name"
//	@router			/sub-municipalities/{subMunicipalityCode}/barangays [get]
func (handler BarangayHandler) ListBySubMunicipalityCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	barangays, err := handler.store.ListBySubMunicipalityCode(vars["subMunicipalityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}

// ListBySpecialGovernmentUnitCode godoc
//
//	@summary		List Barangays
//	@description	List Barangays by Special Government Unit's Philippine Standard Geographic Code (PSGC)
//	@tags			barangays
//	@produce		json
//	@success		200		{object}	stores.Collection[models.Barangay]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			sguCode	path		string	true	"Special Government Unit's PSGC"
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@param			q		query		string	false	"Search by barangay name"
//	@router			/special-government-units/{sguCode}/barangays [get]
func (handler BarangayHandler) ListBySpecialGovernmentUnitCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	opts := NewSearchOptsFromRequest(r)

	barangays, err := handler.store.ListBySpecialGovernmentUnitCode(vars["sguCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}
