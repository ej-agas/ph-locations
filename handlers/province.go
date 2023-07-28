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

func (handler ProvinceHandler) ShowByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	province, err := handler.store.FindByCode(r.Context(), vars["provinceCode"])
	if err != nil {
		JSONResponse(w, ErrProvinceNotFound, http.StatusNotFound)
		return
	}

	JSONResponse(w, province, http.StatusOK)
}

func (handler ProvinceHandler) ListByRegionId(w http.ResponseWriter, r *http.Request) {
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

	provinces, err := handler.store.FindByRegionCode(vars["regionCode"], *opts)

	if err != nil {
		JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, provinces, http.StatusOK)
}
