package routes

import (
	"dewe/handlers"
	"dewe/pkg/middleware"
	"dewe/pkg/mysql"
	"dewe/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	tripRepository := repositories.RepositoryTrips(mysql.DB)
	h := handlers.HandlerTrips(tripRepository)

	r.HandleFunc("/trips", h.FindTrips).Methods("GET")
	r.HandleFunc("/trip/{id}", h.GetTrip).Methods("GET")
	r.HandleFunc("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip))).Methods("POST")
	r.HandleFunc("/trip/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTrip))).Methods("PATCH")
	r.HandleFunc("/trip/{id}", middleware.Auth(h.DeleteTrip)).Methods("DELETE")
}
