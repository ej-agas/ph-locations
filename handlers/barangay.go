package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BarangayHandler struct {
	store stores.BarangayStore
}

func NewBarangayHandler(store stores.BarangayStore) *BarangayHandler {
	return &BarangayHandler{store: store}
}

func (handler BarangayHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	barangay, err := handler.store.FindByCode(vars["barangayCode"])
	if err != nil {
		JSONResponse(w, ErrBarangayNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, barangay, http.StatusOK)
}

func (handler BarangayHandler) ListByCityCode(w http.ResponseWriter, r *http.Request) {
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

	barangays, err := handler.store.ListByCityCode(vars["cityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}

func (handler BarangayHandler) ListByMunicipalityCode(w http.ResponseWriter, r *http.Request) {
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

	barangays, err := handler.store.ListByMunicipalityCode(vars["municipalityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}

func (handler BarangayHandler) ListBySubMunicipalityCode(w http.ResponseWriter, r *http.Request) {
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

	barangays, err := handler.store.ListBySubMunicipalityCode(vars["subMunicipalityCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}

func (handler BarangayHandler) ListBySpecialGovernmentUnitCode(w http.ResponseWriter, r *http.Request) {
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

	barangays, err := handler.store.ListBySpecialGovernmentUnitCode(vars["sguCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, barangays, http.StatusOK)
}
