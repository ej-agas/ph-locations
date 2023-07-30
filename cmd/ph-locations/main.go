package main

import (
	"fmt"
	_ "github.com/ej-agas/ph-locations/docs"
	"github.com/ej-agas/ph-locations/handlers"
	"github.com/ej-agas/ph-locations/postgresql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"os"
	"strconv"
)

//	@title			PSGC API
//	@version		v0.1.0
//	@description	Philippine Standard Geographic Code (PSGC) REST API

//	@contact.name	EJ Agas
//	@contact.url	https://github.com/ej-agas

//	@license.name	GPL v2.0 License
//	@license.url	https://github.com/learning-cloud-native-go/myapp/blob/master/LICENSE

//	@host		localhost:6945
//	@basePath	/api/v1
func main() {
	router := mux.NewRouter()

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal(fmt.Errorf("invalid DB port: %s", err))
	}

	dbConfig := postgresql.Config{
		Host:         os.Getenv("DB_HOST"),
		Port:         port,
		User:         os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName: os.Getenv("POSTGRES_DB"),
	}

	conn, err := postgresql.NewConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	regionStore := postgresql.NewRegionStore(conn)
	regionHandler := handlers.NewRegionHandler(regionStore)

	provinceStore := postgresql.NewProvinceStore(conn)
	provinceHandler := handlers.NewProvinceHandler(provinceStore)

	districtStore := postgresql.NewDistrictStore(conn)
	districtHandler := handlers.NewDistrictHandler(districtStore)

	cityStore := postgresql.NewCityStore(conn)
	cityHandler := handlers.NewCityHandler(cityStore)

	municipalityStore := postgresql.NewMunicipalityStore(conn)
	municipalityHandler := handlers.NewMunicipalityHandler(municipalityStore)

	subMunicipalityStore := postgresql.NewSubMunicipalityStore(conn)
	subMunicipalityHandler := handlers.NewSubMunicipalityHandler(subMunicipalityStore)

	barangayStore := postgresql.NewBarangayStore(conn)
	barangayHandler := handlers.NewBarangayHandler(barangayStore)

	sguStore := postgresql.NewSpecialGovernmentUnit(conn)
	sguHandler := handlers.NewSGUHandler(sguStore)

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	router.HandleFunc("/", handlers.Home)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:6945/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	v1Router := router.PathPrefix("/api/v1/").Subrouter()

	v1Router.HandleFunc("/regions", regionHandler.ListRegions)
	v1Router.HandleFunc("/regions/{regionCode}", regionHandler.ShowRegionByCode)

	v1Router.HandleFunc("/provinces", provinceHandler.List)
	v1Router.HandleFunc("/provinces/{provinceCode}", provinceHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities", cityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities/{cityCode}", cityHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities/{cityCode}/barangays", barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/cities/{cityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities", municipalityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities/{municipalityCode}", municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays", barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units", sguHandler.ListByProvinceCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units/{sguCode}", sguHandler.ShowByCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units/{sguCode}/barangays", barangayHandler.ListBySpecialGovernmentUnitCode)
	v1Router.HandleFunc("/provinces/{provinceCode}/special-government-units/{sguCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/districts", districtHandler.List)
	v1Router.HandleFunc("/districts/{districtCode}", districtHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities", cityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}", cityHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/barangays", barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities", subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities", municipalityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities/{municipalityCode}", municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities/{municipalityCode}/barangays", barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/districts/{districtCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/cities", cityHandler.List)
	v1Router.HandleFunc("/cities/{cityCode}", cityHandler.ShowByCode)
	v1Router.HandleFunc("/cities/{cityCode}/barangays", barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/cities/{cityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities", subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/municipalities", municipalityHandler.List)
	v1Router.HandleFunc("/municipalities/{municipalityCode}", municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/municipalities/{municipalityCode}/barangays", barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/municipalities/{municipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/sub-municipalities", subMunicipalityHandler.List)
	v1Router.HandleFunc("/sub-municipalities/{subMunicipalityCode}", subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/sub-municipalities/{subMunicipalityCode}/barangays/", barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/special-government-units", sguHandler.List)
	v1Router.HandleFunc("/special-government-units/{sguCode}", sguHandler.ShowByCode)
	v1Router.HandleFunc("/special-government-units/{sguCode}/barangays", barangayHandler.ListBySpecialGovernmentUnitCode)
	v1Router.HandleFunc("/special-government-units/{sguCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/barangays", barangayHandler.List)
	v1Router.HandleFunc("/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/districts", districtHandler.ListByRegionId)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}", districtHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities", cityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}", cityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/barangays", barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities", subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities", municipalityHandler.ListByDistrictCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities/{municipalityCode}", municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities/{municipalityCode}/barangays", barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces", provinceHandler.ListByRegionId)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}", provinceHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities", cityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}", cityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/barangays", barangayHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities", municipalityHandler.ListByProvinceCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities/{municipalityCode}", municipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays", barangayHandler.ListByMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/municipalities/{municipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units", sguHandler.ListByProvinceCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units/{sguCode}", sguHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units/{sguCode}/barangays", barangayHandler.ListBySpecialGovernmentUnitCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/special-government-units/{sguCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities", subMunicipalityHandler.ListByCityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}", subMunicipalityHandler.ShowByCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays", barangayHandler.ListBySubMunicipalityCode)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}/cities/{cityCode}/sub-municipalities/{subMunicipalityCode}/barangays/{barangayCode}", barangayHandler.ShowByCode)

	log.Fatal(http.ListenAndServe(":6945", router))
}
