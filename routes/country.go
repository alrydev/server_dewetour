package routes

import (
	"dewe/handlers"
	"dewe/pkg/middleware"
	"dewe/pkg/mysql"
	"dewe/repositories"

	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(countryRepository)

	r.HandleFunc("/countries", middleware.Auth(h.FindCountries)).Methods("GET")
	r.HandleFunc("/country/{id}", middleware.Auth(h.GetCountry)).Methods("GET")
	r.HandleFunc("/country", middleware.Auth(h.CreateCountry)).Methods("POST")
	r.HandleFunc("/country/{id}", middleware.Auth(h.UpdateCountry)).Methods("PATCH")
	r.HandleFunc("/country/{id}", middleware.Auth(h.DeleteCountry)).Methods("DELETE")
}
