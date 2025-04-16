package main

import (
	"net/http"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	//set report caller
	log.SetReportCaller(true)

	//create router
	var router *chi.Mux = chi.NewRouter() 

	//setup routes with handler function
	
	//listen the router
	err := http.ListenAndServe("localhost:8000", router)

	if(err != nil) {
		log.Error(err)
	}
}