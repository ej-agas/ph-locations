package main

import (
	"fmt"
	_ "github.com/ej-agas/ph-locations/docs"
	"github.com/ej-agas/ph-locations/handlers"
	"github.com/ej-agas/ph-locations/postgresql"
	"log"
	"net/http"
	"os"
	"strconv"
)

type config struct {
	port int
	env  string
}

type application struct {
	config                 config
	logger                 *log.Logger
	regionHandler          *handlers.RegionHandler
	provinceHandler        *handlers.ProvinceHandler
	districtHandler        *handlers.DistrictHandler
	cityHandler            *handlers.CityHandler
	municipalityHandler    *handlers.MunicipalityHandler
	subMunicipalityHandler *handlers.SubMunicipalityHandler
	barangayHandler        *handlers.BarangayHandler
	sguHandler             *handlers.SGUHandler
}

var Version string

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
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal(fmt.Errorf("invalid DB port: %s", err))
	}

	cfg := config{
		port: port,
		env:  os.Getenv("APP_ENV"),
	}

	app := &application{
		config: cfg,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
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

	app.regionHandler = handlers.NewRegionHandler(postgresql.NewRegionStore(conn))
	app.provinceHandler = handlers.NewProvinceHandler(postgresql.NewProvinceStore(conn))
	app.districtHandler = handlers.NewDistrictHandler(postgresql.NewDistrictStore(conn))
	app.cityHandler = handlers.NewCityHandler(postgresql.NewCityStore(conn))
	app.municipalityHandler = handlers.NewMunicipalityHandler(postgresql.NewMunicipalityStore(conn))
	app.subMunicipalityHandler = handlers.NewSubMunicipalityHandler(postgresql.NewSubMunicipalityStore(conn))
	app.barangayHandler = handlers.NewBarangayHandler(postgresql.NewBarangayStore(conn))
	app.sguHandler = handlers.NewSGUHandler(postgresql.NewSpecialGovernmentUnit(conn))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), app.routes()))
}
