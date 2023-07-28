package main

import (
	"fmt"
	"github.com/ej-agas/ph-locations/handlers"
	"github.com/ej-agas/ph-locations/postgresql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

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

	provinceStore := postgresql.NewProvinceStore(conn)
	provinceHandler := handlers.NewProvinceHandler(provinceStore)

	districtStore := postgresql.NewDistrictStore(conn)
	districtHandler := handlers.NewDistrictHandler(districtStore)

	regionStore := postgresql.NewRegionStore(conn)
	regionHandler := handlers.NewRegionHandler(regionStore)

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	router.HandleFunc("/", handlers.Home)
	v1Router := router.PathPrefix("/api/v1/").Subrouter()
	v1Router.HandleFunc("/provinces/{id}", provinceHandler.ShowProvinceById)
	v1Router.HandleFunc("/provinces/{code}", provinceHandler.ShowProvinceById)

	v1Router.HandleFunc("/regions", regionHandler.ListRegions)
	v1Router.HandleFunc("/regions/{regionCode}", regionHandler.ShowRegionByCode)

	v1Router.HandleFunc("/regions/{regionCode}/provinces", provinceHandler.ListByRegionId)
	v1Router.HandleFunc("/regions/{regionCode}/provinces/{provinceCode}", provinceHandler.ShowByCode)

	v1Router.HandleFunc("/regions/{regionCode}/districts", districtHandler.ListByRegionId)
	v1Router.HandleFunc("/regions/{regionCode}/districts/{districtCode}", districtHandler.ShowByCode)

	log.Fatal(http.ListenAndServe(":6945", router))
}
