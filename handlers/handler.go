package handlers

import (
	"github.com/go-chi/chi"
	"superheroapi/middleware"
)


func Handler(router *chi.Mux) {

	//use go-chi to strip slashes

	//route
	router.Route("/superhero", func (route chi.Router) {
		//use middleware to authenticate with route.use
		route.Use(middleware.Authorize)
		//grab data
		route.Get("/ability", GetSuperheroAbility)
	})
}