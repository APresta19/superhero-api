package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"superheroapi/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	//set report caller
	log.SetReportCaller(true)

	//create router
	var router *chi.Mux = chi.NewRouter() 

	//setup routes with handler function
	handlers.Handler(router)
	
	fmt.Println("Starting Superhero API...")

	//listen the router
	err := http.ListenAndServe("localhost:8000", router)

	if(err != nil) {
		log.Error(err)
		return
	}

	fmt.Println("Server started!")
}