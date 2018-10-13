package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../models"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var configuration server.PersistenceServiceConfiguration
var endpoint *server.PersistenceServiceEndpoint

func main() {
	logging.CreateLog()
	logging.LogHeader("Persistence Service")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)
	endpoint = server.NewPersistenceServiceEndpoint(configuration)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetKeyValueCache", getKeyValueCache).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/SetKeyValueCache", setKeyValueCache).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/GetInfrastructureMetadata", getInfrastructureMetadata).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/SetLogRecord", setLogRecord).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func getKeyValueCache(w http.ResponseWriter, r *http.Request) {
	var model models.KeyValueGetRequest
	err := jsonutil.DecodeJsonFromBody(r, &model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	result, err := endpoint.GetKeyValueCache()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	resultBytes, err := jsonutil.EncodeJsonToBytes(&result)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	w.Write(*resultBytes)
}

func setKeyValueCache(w http.ResponseWriter, r *http.Request) {
	/*
		var model models.KeyValueSetRequest
		err := jsonutil.DecodeJsonFromBody(r, &model)
		if err != nil {
			w.WriteHeader(500)
			logging.LogError(err)
			return
		}
		result, err := endpoint.something(something)
		if err != nil {
			w.WriteHeader(500)
			logging.LogError(err)
			return
		}
		resultBytes, err := jsonutil.EncodeJsonToBytes(&result)
		if err != nil {
			w.WriteHeader(500)
			logging.LogError(err)
			return
		}
		w.Write(*resultBytes)
	*/
}

func getInfrastructureMetadata(w http.ResponseWriter, r *http.Request) {
	var model models.InfrastructureMetadata
	err := jsonutil.DecodeJsonFromBody(r, &model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	result, err := endpoint.GetInfrastructureMetadata()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	resultBytes, err := jsonutil.EncodeJsonToBytes(&result)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	w.Write(*resultBytes)
}

func setLogRecord(w http.ResponseWriter, r *http.Request) {
	/*
		var model models.Model
		err := jsonutil.DecodeJsonFromBody(r, &model)
		if err != nil {
			w.WriteHeader(500)
			logging.LogError(err)
			return
		}
		result, err := endpoint.something(something)
		if err != nil {
			w.WriteHeader(500)
			logging.LogError(err)
			return
		}
		resultBytes, err := jsonutil.EncodeJsonToBytes(&result)
		if err != nil {
			w.WriteHeader(500)
			logging.LogError(err)
			return
		}
		w.Write(*resultBytes)
	*/
}
