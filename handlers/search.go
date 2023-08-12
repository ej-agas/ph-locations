package handlers

import (
	"github.com/ej-agas/ph-locations/stores"
	"net/http"
	"strconv"
)

func NewSearchOptsFromRequest(r *http.Request) *stores.SearchOpts {
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

	search := r.URL.Query().Get("q")

	return stores.NewSearchOpts(
		stores.WithSort(sort),
		stores.WithOrder(order),
		stores.WithLimit(limit),
		stores.WithPage(page),
		stores.WithSearch(search),
	)
}
