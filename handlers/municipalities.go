package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MunicipalityHandler struct {
	store stores.MunicipalityStore
}

func NewMunicipalityHandler(store stores.MunicipalityStore) *MunicipalityHandler {
	return &MunicipalityHandler{store: store}
}

func (handler MunicipalityHandler) FindByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	municipality, err := handler.store.FindByCode(vars["municipalityCode"])
	if err != nil {
		JSONResponse(w, ErrMunicipalityNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, municipality, http.StatusOK)
}

func (handler MunicipalityHandler) ListByProvinceCode(w http.ResponseWriter, r *http.Request) {
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

	provinces, err := handler.store.ListByProvinceCode(vars["provinceCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}

func (handler MunicipalityHandler) ListByDistrictCode(w http.ResponseWriter, r *http.Request) {
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

	provinces, err := handler.store.ListByDistrictCode(vars["districtCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}
