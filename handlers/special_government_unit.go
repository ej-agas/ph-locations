package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type SGUHandler struct {
	store stores.SpecialGovernmentUnit
}

func NewSGUHandler(store stores.SpecialGovernmentUnit) *SGUHandler {
	return &SGUHandler{store: store}
}

func (handler SGUHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sgu, err := handler.store.FindByCode(vars["sguCode"])
	if err != nil {
		JSONResponse(w, ErrSGUNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, sgu, http.StatusOK)
}

func (handler SGUHandler) ListByProvinceCode(w http.ResponseWriter, r *http.Request) {
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

	sgus, err := handler.store.ListByProvinceCode(vars["provinceCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, sgus, http.StatusOK)
}
