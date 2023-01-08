package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	CountryRoutes(r)
	UserRoutes(r)
	AuthRoutes(r)
	TripRoutes(r)
	TransactionRoutes(r)
}
