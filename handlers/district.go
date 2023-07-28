package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DistrictHandler struct {
	store stores.DistrictStore
}

func NewDistrictHandler(store stores.DistrictStore) *DistrictHandler {
	return &DistrictHandler{store: store}
}

func (handler DistrictHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	district, err := handler.store.FindByCode(vars["districtCode"])
	if err != nil {
		JSONResponse(w, ErrDistrictNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, district, http.StatusOK)
}

func (handler DistrictHandler) ListByRegionId(w http.ResponseWriter, r *http.Request) {
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

	districts, err := handler.store.FindByRegionCode(vars["regionCode"], *opts)
	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, districts, http.StatusOK)
}
