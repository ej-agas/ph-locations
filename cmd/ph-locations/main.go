package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ph-locations/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Home)

	log.Fatal(http.ListenAndServe(":1337", router))
}
