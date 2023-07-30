package handlers

import (
	"fmt"
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
)

type RegionHandler struct {
	store stores.RegionStore
}

func NewRegionHandler(store stores.RegionStore) *RegionHandler {
	return &RegionHandler{store: store}
}

// List godoc
//
//	@summary		List Regions
//	@description	List Regions
//	@tags			regions
//	@produce		json
//	@success		200		{object}	stores.Collection[models.Region]
//	@failure		404		{object}	handlers.ResponseMessage
//	@param			order	query		string	false	"Order by id, code (PSGC), Name, Population. (default: id)"
//	@param			sort	query		string	false	"Sort by asc (Ascending) desc (Descending). (default: asc)"
//	@param			limit	query		string	false	"Limit results per page. (default: 25)"
//	@param			page	query		string	false	"Page number. (default: 1)"
//	@router			/regions [get]
func (handler RegionHandler) List(w http.ResponseWriter, r *http.Request) {
	opts := NewSearchOptsFromRequest(r)

	regions, err := handler.store.List(*opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, regions, http.StatusOK)
}

// ShowRegionByCode godoc
//
//	@summary		Show Region
//	@description	Show Region by Philippine Standard Geographic Code (PSGC)
//	@tags			regions
//	@produce		json
//	@success		200			{object}	models.Region
//	@failure		404			{object}	handlers.ResponseMessage
//	@param			regionCode	path		string	true	"Region PSGC"
//	@router			/regions/{regionCode} [get]
func (handler RegionHandler) ShowRegionByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["regionCode"]
	fmt.Println(code)
	region, err := handler.store.FindByCode(code)
	fmt.Println(region, err)
	if err != nil {
		JSONResponse(w, ErrRegionNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, region, http.StatusOK)
}
