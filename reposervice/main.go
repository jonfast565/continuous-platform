package main

import (
	"../utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var configuration TeamServicesConfiguration
var endpoint *TeamServicesEndpoint

func main() {

	utilities.LogApplicationStart()
	file, err := utilities.CreateLog()
	if err != nil {
		panic("Log not created")
	}
	defer file.Close()

	utilities.DecodeJsonFromFile("./appsettings.json", configuration)
	endpoint = NewTeamServicesEndpoint(configuration)

	router := mux.NewRouter()
	router.HandleFunc("/repositories", getRepositories).Methods(utilities.PostMethod)
	router.HandleFunc("/file", getFile).Methods(utilities.PostMethod)

	log.Print("Serving content...")
	log.Fatal(http.ListenAndServe(utilities.GetLocalPort(configuration.Port), router))
	utilities.LogApplicationEnd()
}

func getRepositories(w http.ResponseWriter, r *http.Request) {
	// endpoint.
}

func getFile(w http.ResponseWriter, r *http.Request) {
	// endpoint.
}
