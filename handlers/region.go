package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ph-locations/stores"
)

type RegionHandler struct {
	store stores.RegionStore
}

func NewRegionHandler(store stores.RegionStore) *RegionHandler {
	return &RegionHandler{store: store}
}

func (handler RegionHandler) ListRegions(w http.ResponseWriter, r *http.Request) {
	regions, err := handler.store.All()

	if err != nil {
		JSONResponse(w, ResponseMessage{StatusCode: 500, Message: err.Error()}, 500)
		return
	}

	JSONResponse(w, regions, 200)
}

func (handler RegionHandler) ShowRegionByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Println(code)
	region, err := handler.store.FindByCode(code)
	fmt.Println(region, err)
	if err != nil {
		JSONResponse(w, ResponseMessage{StatusCode: 404, Message: "region not found"}, 404)
		return
	}

	JSONResponse(w, region, 200)
}
