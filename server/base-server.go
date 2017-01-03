package server

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

const base_path = "/api/v1/"
var router *mux.Router

func SetupRestServer() {
	// Register a handler for each route pattern
	router = mux.NewRouter()

	// Add a trivial handler for INFO
	router.Methods("GET").Path(base_path + "animals").HandlerFunc(HandleGetAnimals)
}

func StartRestServer() {
	// Get Cloud Foundry assigned port
	port := "8080"

	// Start listening on the configured port.
	err:= http.ListenAndServe(":"+port, router)
	if(err != nil){
		os.Exit(1)
	}
}
