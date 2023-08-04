package main

import (
	"fmt"
	"github.com/ej-agas/ph-locations/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
	"os"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	router.HandleFunc("/", handlers.Home)
	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/docs/doc.json", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	v1Router := router.PathPrefix("/api/v1/").Subrouter()

	v1Router.HandleFunc("/regions", app.regionHandler.List)
	v1Router.HandleFunc("/regions/{regionCode}", app.regionHandler.ShowByCode)

	v1Router.HandleFunc("/provinces", app.provinceHandler.List)
	v1Router.HandleFunc("/provinces/{provinceCode}", app.provinceHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities", app.cityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities/{cityCode}", app.cityHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities/{cityCode}/barangays", app.barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities/{cityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities", app.municipalityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities/{municipalityCode}", app.municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays", app.barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units", app.sguHandler.ListByProvinceCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units/{sguCode}", app.sguHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units/{sguCode}/barangays", app.barangayHandler.ListBySpecialGovernmentUnitCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units/{sguCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/districts", app.districtHandler.List)
	v1Router.HandleFunc("/districts/{districtCode}", app.districtHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities", app.cityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}", app.cityHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/barangays", app.barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities", app.subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", app.subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", app.barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities", app.municipalityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities/{municipalityCode}", app.municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities/{municipalityCode}/barangays", app.barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/cities", app.cityHandler.List)
	v1Router.HandleFunc("/cities/{cityCode}", app.cityHandler.ShowByCode)
	v1Router.HandleFunc("/cities/{cityCode}/barangays", app.barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/cities/{cityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities", app.subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", app.subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", app.barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/municipalities", app.municipalityHandler.List)
	v1Router.HandleFunc("/municipalities/{municipalityCode}", app.municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/municipalities/{municipalityCode}/barangays", app.barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/municipalities/{municipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/sub-municipalities", app.subMunicipalityHandler.List)
	v1Router.HandleFunc("/sub-municipalities/{subMunicipalityCode}", app.subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/sub-municipalities/{subMunicipalityCode}/barangays/", app.barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/special-government-units", app.sguHandler.List)
	v1Router.HandleFunc("/special-government-units/{sguCode}", app.sguHandler.ShowByCode)
	v1Router.HandleFunc("/special-government-units/{sguCode}/barangays", app.barangayHandler.ListBySpecialGovernmentUnitCode)
	v1Router.HandleFunc("/special-government-units/{sguCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/barangays", app.barangayHandler.List)
	v1Router.HandleFunc("/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/districts", app.districtHandler.ListByRegionCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}", app.districtHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities", app.cityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}", app.cityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/barangays", app.barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities", app.subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", app.subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", app.barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities", app.municipalityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities/{municipalityCode}", app.municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities/{municipalityCode}/barangays", app.barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces", app.provinceHandler.ListByRegionCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}", app.provinceHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities", app.cityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}", app.cityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/barangays", app.barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities", app.municipalityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities/{municipalityCode}", app.municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays", app.barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units", app.sguHandler.ListByProvinceCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units/{sguCode}", app.sguHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units/{sguCode}/barangays", app.barangayHandler.ListBySpecialGovernmentUnitCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units/{sguCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities", app.subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", app.subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", app.barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", app.barangayHandler.ShowByCode)

	return router
}
