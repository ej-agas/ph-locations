package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type SubMunicipalityHandler struct {
	store stores.SubMunicipality
}

func NewSubMunicipalityHandler(store stores.SubMunicipality) *SubMunicipalityHandler {
	return &SubMunicipalityHandler{store: store}
}

func (handler SubMunicipalityHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city, err := handler.store.FindByCode(vars["subMunicipalityCode"])
	if err != nil {
		JSONResponse(w, ErrSubMunicipalityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, city, http.StatusOK)
}

func (handler SubMunicipalityHandler) ListByCityCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	allowedColumns := []string{"id", "code", "name", "population"}

	sort := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 25
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	if IsInAllowedColumns(order, allowedColumns) == false {
		order = "id"
	}

	opts := stores.NewSearchOpts(
		stores.WithSort(sort),
		stores.WithOrder(order),
		stores.WithLimit(limit),
		stores.WithPage(page),
	)

	subMunicipalities, err := handler.store.ListByCityCode(vars["cityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, subMunicipalities, http.StatusOK)
}
